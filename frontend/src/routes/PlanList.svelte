<script>
  import { onMount } from "svelte";
  import { GetAllPlans, SaveConfig } from "$lib/wailsjs/go/main/App";

  export let modalOpen = false;
  export let config;
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
                config.CurrentPlan = plan.Name;
                SaveConfig(config);
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
