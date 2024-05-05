<script>
  import { onMount } from "svelte";
  import { GetAllPlans, SaveConfig } from "$lib/wailsjs/go/main/App";
  import { config } from "./stores.js";

  let config_value;
  config.subscribe((c) => {
    config_value = c;
  });

  export let modalOpen = false;

  let all_plans;

  onMount(() => {
    GetAllPlans().then((res) => {
      all_plans = res;
    });
  });
</script>

<dialog open={modalOpen}>
  <article>
    <header>
      <button aria-label="close" rel="prev" on:click={() => (modalOpen = false)}
      ></button>
      <p><strong>Choose plan</strong></p>
    </header>

    {#if all_plans}
      <div class="overflow-auto">
        {#each all_plans as plan}
          <details>
            <summary>{plan.Name}</summary>
            <p>{plan.Description}</p>
            <button
              on:click={() => {
                config_value.CurrentPlan = plan.Name;
                SaveConfig(config_value);
                config.set(config_value);
                modalOpen = false;
              }}>Select</button
            >
          </details>
        {/each}
      </div>
    {:else}
      <p>No plans found</p>
    {/if}
  </article>
</dialog>
