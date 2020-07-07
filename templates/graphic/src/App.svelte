<script>
  import * as d3 from "d3";
  import legend from "d3-svg-legend";
  import { onMount } from "svelte";
  export let used,total
  let svgElement;
  onMount(() => {
    let dims = {
      width: 250,
      height: 250,
      margins: 20
    };
    let svg = d3.select(svgElement);
    svg.attr("height", dims.height);
    svg.attr("width", dims.width);
    let g = svg
      .append("g")
      .attr(
        "transform",
        `translate(${dims.margins + dims.width / 2},${dims.margins +
          dims.height / 2})`
      );
    let pie = d3.pie()([used, total]);
    console.log(pie);
    let arcgen = d3
      .arc()
      .innerRadius(70)
      .outerRadius(100);
    g.selectAll(".graph")
      .data(pie)
      .enter()
      .append("path")
      .attr("class", "graph")
      .attr("fill", (d, i) => {
        if (i == 0) {
          // purplish
          return "black";
        }
        return "white";
      })
      .attr("stroke", "black")
      .attr("d", arcgen);
    g.selectAll("text")
      .data(pie)
      .enter()
      .append("text")
      .attr("transform", d => `translate(${arcgen.centroid(d)})`)
      .style("text-anchor", "middle")
      .text(d => `${d.data}TB`)
      .attr("fill", (d, i) => {
        if (i == 1) {
          // purplish
          return "black";
        }
        return "white";
	  })
	  .attr("class","storageText")
  });
</script>

<div id="innerGraphicHolder">
  <svg id="holder" bind:this={svgElement} />
</div>
