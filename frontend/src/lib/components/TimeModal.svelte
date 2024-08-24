<script>
  export let modalOpen = false;
  export let remaining_time = 0;
  export let max_time = 0;
</script>

<dialog open={modalOpen} class="time-modal">
  <article>
    <header>Rest</header>
    {new Date(remaining_time).toISOString().substring(11, 19)}
    <progress value={remaining_time} max={max_time} />
  </article>
</dialog>

<style>
  .time-modal {
    width: 20%;
    position: fixed;
    bottom: 0;
    right: 0;
  }
</style>
