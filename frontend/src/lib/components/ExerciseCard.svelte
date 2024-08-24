<script>
  export let set_idx;
  export let set;
  export let rest_time;
  import Icon from "./Icon.svelte";
  import { rest_timer } from "../../routes/stores";

  let is_open = true;

  let toggle_input = () => {
    document.getElementById(`RepInput${set_idx}`).disabled =
      !document.getElementById(`RepInput${set_idx}`).disabled;
    document.getElementById(`WeightInput${set_idx}`).disabled =
      !document.getElementById(`WeightInput${set_idx}`).disabled;
    is_open = !is_open;

    if (!is_open) {
      is_break = true;
      rest_timer.set(rest_time);
    }
  };
</script>

<details open={is_open}>
  <summary class="center">Set {set_idx}</summary>
  <div role="group">
    Reps:
    <button
      on:click={() => {
        set.Repetitions -= 1;
      }}><Icon name="minus" /></button
    >
    <input
      type="number"
      min="1"
      step="1"
      bind:value={set.Repetitions}
      id={`RepInput${set_idx}`}
    />
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
    <input
      type="number"
      min="0"
      step="0.25"
      bind:value={set.Weight}
      id={`WeightInput${set_idx}`}
    />
    <button
      on:click={() => {
        set.Weight += 1;
      }}><Icon name="plus" /></button
    >
  </div>
  <div role="group">
    <button on:click={toggle_input} class={is_open ? "primary" : "secondary"}
      >{is_open ? "Done" : "Edit"}</button
    >
  </div>
</details>

<style>
  .center {
    text-align: center;
    display: flex;
    flex-flow: column;
    justify-content: center;
  }
</style>
