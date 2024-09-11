<script>
  import { onMount } from "svelte";
  import Icon from "./Icon.svelte";
  import { rest_timer } from "../../routes/stores";
  export let is_break;
  export let max_time;
  let dialog = null;
  let remaining_time;

  onMount(() => {
    dialog = document.getElementById("time-dialog");
  });

  let convertToMinutesSeconds = (s) => {
    s = Math.floor(s);
    return (s - (s %= 60)) / 60 + (9 < s ? ":" : ":0") + s;
  };

  rest_timer.subscribe((rt) => {
    remaining_time = rt;
  });

  $: {
    if (is_break && dialog !== null) {
      dialog.style.setProperty("z-index", "1000");
      dialog.style.setProperty("right", "10px");
    } else if (dialog !== null) {
      dialog.style.setProperty("z-index", "-9999");
      dialog.style.setProperty("right", "-9999px");
    }
  }
</script>

<div class="time-modal" id="time-dialog">
  <article data-theme="light">
    <header>
      <button on:click={() => rest_timer.set(0)}> Stop timer </button>
    </header>
    <p style="color: black;">Rest: {convertToMinutesSeconds(remaining_time)}</p>
    <progress value={remaining_time} max={max_time} />
  </article>
</div>

<style>
  .time-modal {
    width: 20%;
    position: fixed;
    right: 10px;
    bottom: 0;
  }

  @media (max-width: 768px) {
    .time-modal {
      width: 50%;
    }
  }
</style>
