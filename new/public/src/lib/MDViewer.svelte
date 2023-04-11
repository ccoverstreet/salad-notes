<script>
	export let content;

	let display = {};
	$: display.innerHTML = content;
</script>

<div id="output" bind:this={display}></div>

<style>
</style>

