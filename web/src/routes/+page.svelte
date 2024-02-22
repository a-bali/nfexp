<script>
    import { Chart } from "svelte-echarts";
    import { formatBytes } from "$lib/format.js";
    export let data;
    var options;

    data.overview.promise.then((data) => {
        var chartdata = [];
        var lines = data.split("\n");
        lines.slice(1).forEach((line) => {
            if (line.length == 0) return;
            var fields = line.split(",");
            chartdata.push([fields[0], fields[1], fields[3]]);
        });

        options = {
            tooltip: {
                trigger: "axis",
                axisPointer: { type: "cross" },
            },
            toolbox: {
                feature: {
                    dataZoom: {
                        yAxisIndex: false,
                    },
                    restore: {},
                },
            },
            dataZoom: [
                {
                    type: "inside",
                },
                {
                    start: 0,
                    end: 100,
                },
            ],
            dataset: {
                source: chartdata,
                dimensions: ["timestamp", "flows", "bytes"],
            },
            xAxis: { type: "time" },
            yAxis: [
                { name: "flows", type: "value" },
                {
                    name: "bytes",
                    type: "value",
                    axisLabel: {
                        formatter: formatBytes,
                    },
                },
            ],
            series: [
                {
                    name: "bytes",
                    type: "bar",
                    yAxisIndex: 1,
                    encode: {
                        x: "timestamp",
                        y: "bytes",
                    },
                },
                {
                    name: "flows",
                    type: "line",
                    smooth: true,
                    yAxisIndex: 0,
                    encode: {
                        x: "timestamp",
                        y: "flows",
                    },
                },
            ],
        };
    });
</script>

<h1>Overview</h1>
<main>
    {#await data.overview.promise}
        (loading)
    {:then}
        <div class="app">
            <Chart {options} />
        </div>
    {:catch error}
        Error: {error.message}
    {/await}
</main>

<style>
    .app {
        width: 100vw;
        height: 80vh;
    }
</style>
