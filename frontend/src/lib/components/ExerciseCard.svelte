<script>
  export let set_idx;
  export let set;
  import Icon from "./Icon.svelte";
</script>

<article>
  <header class="center">Set {set_idx}</header>
  <div role="group">
    Reps:
    <Icon name="minus" />
    {set.Repetitions}
    <Icon name="plus" />
  </div>
  <div role="group">
    Weight:
    <Icon name="minus" />
    {set.Weight} kg
    <Icon name="plus" />
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
