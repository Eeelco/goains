<script>
  import { onMount } from "svelte";
  import { GetAllPlans, SetCurrentPlan } from "../lib/wailsjs/go/main/App";

  export let modalOpen = false;
  export let plan_name = "";
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
      <h1>Choose plan</h1>
      <button on:click={() => (modalOpen = false)}>Close</button>
    </header>

    {#if all_plans}
      <div class="overflow-auto">
        {#each all_plans as plan}
          <details>
            <summary>{plan.Name}</summary>
            <p>{plan.Description}</p>
            <button
              on:click={() => {
                SetCurrentPlan(plan.Name);
                plan_name = plan.Name;
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
