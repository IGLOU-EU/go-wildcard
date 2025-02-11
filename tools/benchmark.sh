#!/bin/bash

go_bench_cmd="go test -benchmem -bench . github.com/IGLOU-EU/go-wildcard/v2/benchmark"

assets="./assets"
result_raw="${assets}/result_raw.txt"
graph_time="${assets}/graph_time.png"
graph_mem="${assets}/graph_allocs.png"

mkdir -p "${assets}"
true > "${result_raw}"

bench_name=()
bench_set=()
bench_ns=()
bench_allocs=()

while read -r line ; do
    echo "$line" >> "${result_raw}"
    _name=$(echo "$line" | awk '{print $1}')

    bench_name+=("${_name%/*}")
    bench_set+=("${_name#*/}")
    bench_ns+=("$(echo "$line" | awk '{print $3}')")
    bench_allocs+=("$(echo "$line" | awk '{print $7}')")
done < <(bash -c "$go_bench_cmd" | grep "allocs/op")


unique_names=()
for ((i=0; i<${#bench_name[@]}; i++)); do
    name="${bench_name[i]}"
    found=0
    for existing in "${unique_names[@]}"; do
        if [[ "$existing" == "$name" ]]; then
            found=1
            break
        fi
    done
    if [[ $found -eq 0 ]]; then
        unique_names+=("$name")
    fi
done

declare -A bench_colors
color_palette=("red" "blue" "green" "orange" "magenta" "cyan" "brown" "purple")

for idx in "${!unique_names[@]}"; do
    bench_colors["${unique_names[idx]}"]="${color_palette[idx % ${#color_palette[@]}]}"
done

generate_plot() {
    local data_array=$1
    local title="$2"
    local ylabel="$3"
    local max_value=$4 
    local output_file=$5

    for name in "${unique_names[@]}"; do
        data_file="/tmp/bench_data_${name}_${title// /_}.dat"
        true > "$data_file"
        for ((i=0; i<${#bench_name[@]}; i++)); do
            if [[ "${bench_name[i]}" == "$name" ]]; then
                x="${bench_set[i]}"
                eval "y=\${${data_array}[i]}"
                truncate=$(awk -v y="$y" -v max="$max_value" 'BEGIN { if (y > max) print "1"; else print "0" }')
                if [ "$truncate" -eq 1 ]; then
                    y=$max_value
                fi
                echo "$x $y" >> "$data_file"
            fi
        done
    done

    plot_cmd="/tmp/plot_bench_${title// /_}.gnuplot"
    cat <<EOF > "$plot_cmd"
set terminal png enhanced size 1200,400 font '/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf,10'
set output "$output_file"
set size 1,1
set title "$title"
set xlabel "bench_set"
set ylabel "$ylabel"
set key below center horizontal
# Keep automatic tics and add a special label for max value
set ytics auto
set ytics add ("$max_value or more" $max_value)
EOF

    plot_line="plot "
    first=1
    for name in "${unique_names[@]}"; do
        data_file="/tmp/bench_data_${name}_${title// /_}.dat"
        color="${bench_colors[$name]}"
        if [ $first -eq 1 ]; then
            first=0
        else
            plot_line+=", "
        fi
        plot_line+="\"$data_file\" using 1:2 with linespoints linecolor rgb \"$color\" title \"$name\""
    done

    echo "$plot_line" >> "$plot_cmd"

    gnuplot "$plot_cmd"
}

generate_plot bench_ns "Benchmark Time" "bench_ns (ns)" 300 "$graph_time" 
generate_plot bench_allocs "Benchmark Allocations" "bench_allocs (allocs/op)" 10 "$graph_mem"
