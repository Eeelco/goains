<script>
  import { SaveConfig } from "$lib/wailsjs/go/main/App";

  export let settingsOpen = false;
  export let config;
</script>

<dialog open={settingsOpen}>
  <article>
    <header>
      <button
        aria-label="close"
        rel="prev"
        on:click={() => (settingsOpen = false)}
      ></button>
      <p><strong>Settings</strong></p>
    </header>
    <h2>Default values</h2>
    <form>
      <label
        >Default reps
        <input type="number" bind:value={config.DefaultNrReps} />
      </label>
      <label for="default_sets">Default sets</label>
      <input
        type="number"
        id="default_sets"
        bind:value={config.DefaultNrSets}
      />
      <label for="default_rest">Default rest</label>
      <input type="number" id="default_rest" bind:value={config.DefaultRest} />
      <button
        type="submit"
        on:click={() => {
          SaveConfig(config);
          settingsOpen = false;
        }}>Save</button
      >
    </form>
  </article>
</dialog>
