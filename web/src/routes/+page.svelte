<script>
    import { Chart } from "svelte-echarts";
    import { formatBytes } from "$lib";
    export let data;
    var options;
    var nodata;

    $: data, setupChart();

    let setupChart = () => {
        data.overview.promise.then((data) => {
            var chartdata = [];
            nodata = true;
            var lines = data.split("\n");
            lines.slice(1).forEach((line) => {
                if (line.length == 0) return;
                var fields = line.split(",");
                chartdata.push([fields[0], fields[1], fields[3]]);
                nodata = false;
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
    };
</script>

<h1>Traffic / flow overview</h1>
<main>
    {#await data.overview.promise}
        (loading)
    {:then}
        {#if nodata}
            <div>No data</div>
        {:else}
            <div class="app">
                <Chart {options} />
            </div>
        {/if}
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
