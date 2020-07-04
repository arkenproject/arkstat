<script>
import * as d3 from "d3"
import legend from "d3-svg-legend"
import {onMount} from "svelte"
let svgElement
onMount(()=> {
	let dims = {
		width:250,
		height:250,
		margins:20,
	} 
	let svg = d3.select(svgElement)
	svg.attr("height",dims.height)
	svg.attr("width",dims.width)
	let g = svg.append("g").attr("transform",`translate(${dims.margins + dims.width/2},${dims.margins + dims.height/2})`)
	let used = 50
	let total = 500
	let pie = d3.pie()([used,total])
	g.selectAll(".graph").data(pie)
	.enter()
	.append("path")
	.attr("class","graph")
	.attr("fill",(d,i)=> {
		if (i == 0) {
			// purplish
			return "#693AFA"
		}
		return "white"
	})
	.attr("stroke","black")
	.attr("d",d3.arc()
	.innerRadius(80)
	.outerRadius(100))

	

})

</script>

<div id="innerGraphicHolder">
<svg id="holder" bind:this={svgElement}></svg>
</div>