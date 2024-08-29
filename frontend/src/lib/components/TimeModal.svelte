<script>
  import { onMount } from "svelte";
  import Icon from "./Icon.svelte";
  import { rest_timer } from "../../routes/stores";
  export let modalOpen;
  export let remaining_time;
  export let max_time;
  let dialog = null;
  let interval = null;

  onMount(() => {
    dialog = document.getElementById("time-dialog");
  });

  let fmtMSS = (s) => {
    s = Math.floor(s);  
    return(s-(s%=60))/60+(9<s?':':':0')+s
  }

  $: {
    if (modalOpen && dialog !== null) {
      dialog.style.setProperty("z-index", "1000");
      dialog.style.setProperty("right", "10px");
    } else if (dialog !== null) {
      dialog.style.setProperty("z-index", "-9999");
      dialog.style.setProperty("right", "-9999px");
    }
  }

</script>

<div class="time-modal" id="time-dialog">
  <article>
    <header>
      Rest
      <button on:click={() => (rest_timer.set(0))}>
        <Icon name="close" />
        </button>
    </header>
     {fmtMSS(remaining_time)}
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
</style>
