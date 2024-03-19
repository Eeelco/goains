<script>
  import { onMount } from "svelte";

  export let modalOpen = false;

  function openModal() {
    modalOpen = true;
  }

  onMount(() => {
    // Initialize PicoCSS modal here
  });
</script>

<dialog open={modalOpen}>
  <article>
    <h1>Modal</h1>
    <button on:click={() => (modalOpen = false)}>Close</button>

    <p>Modal content</p>
  </article>
</dialog>
