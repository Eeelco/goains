<script>
  export let set_idx;
  export let set;
  import Icon from "./Icon.svelte";
</script>

<article>
  <header class="center">Set {set_idx}</header>
  <div role="group">
    Reps:
    <button
      on:click={() => {
        set.Repetitions -= 1;
      }}><Icon name="minus" /></button
    >
    <input type="number" min="1" step="1" bind:value={set.Repetitions} />
    <button
      on:click={() => {
        set.Repetitions += 1;
      }}><Icon name="plus" /></button
    >
  </div>
  <div role="group">
    Weight:
    <button
      on:click={() => {
        set.Weight -= 1;
      }}><Icon name="minus" /></button
    >
    <input type="number" min="0" step="0.25" bind:value={set.Weight} />
    <button
      on:click={() => {
        set.Weight += 1;
      }}><Icon name="plus" /></button
    >
  </div>
  <footer></footer>
</article>

<style>
  .center {
    text-align: center;
    display: flex;
    flex-flow: column;
    justify-content: center;
  }
</style>
