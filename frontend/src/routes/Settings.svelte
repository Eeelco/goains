<script>
  import { SaveConfig } from "$lib/wailsjs/go/main/App";
  import { config } from "./stores.js";

  let config_value;
  config.subscribe((c) => {
    config_value = c;
  });

  export let settingsOpen = false;
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
        <input type="number" bind:value={config_value.DefaultNrReps} />
      </label>
      <label for="default_sets">Default sets</label>
      <input
        type="number"
        id="default_sets"
        bind:value={config_value.DefaultNrSets}
      />
      <label for="default_rest">Default rest</label>
      <input
        type="number"
        id="default_rest"
        bind:value={config_value.DefaultRest}
      />
      <button
        type="submit"
        on:click={() => {
          SaveConfig(config_value);
          config.set(config_value);
          settingsOpen = false;
        }}>Save</button
      >
    </form>
  </article>
</dialog>
