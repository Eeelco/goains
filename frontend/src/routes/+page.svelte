<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { GetConfig, GetPlan, StartWorkout } from "$lib/wailsjs/go/backend/App";
  import Icon from "$lib/components/Icon.svelte";
  import { config, current_day_idx, plan } from "./stores.js";

  import PlanList from "./PlanList.svelte";
  import Settings from "./Settings.svelte";

  const rotateArray = (arr, k) =>
    arr ? arr.slice(k).concat(arr.slice(0, k)) : [];
  const LoadPlan = () => {
    GetPlan(config_value.CurrentPlan).then((res) => {
      plan.set(res);
      days = rotateArray(plan_value.Days, config_value.NextDayIdx);
    });
  };

  let config_value;
  let plan_value;

  // On startup, config is loaded and then plan is loaded
  // and set to the current day
  config.subscribe((c) => {
    config_value = c;
    LoadPlan();
  });
  plan.subscribe((p) => {
    plan_value = p;
  });

  let modalOpen = false;
  let settingsOpen = false;

  let days = [];
  onMount(() => {
    GetConfig().then((res) => {
      config.set(res);
    });
  });
</script>

<PlanList bind:modalOpen />
<Settings bind:settingsOpen />
<nav>
  <ul>
    <li>
      <button
        class="outline"
        on:click={() => {
          settingsOpen = true;
        }}
      >
        <Icon name="cog" />
      </button>
    </li>
  </ul>
</nav>
<div class="center">
  <h1>Goains</h1>

  {#if config_value.CurrentPlan !== ""}
    <details open>
      <summary role="button" class="secondary"
        >Continue {config_value.CurrentPlan}</summary
      >
      {#each days as day, idx}
        <button
          data-sveltekit-preload-data="hover"
          class="breathe"
          on:click={() => {
            current_day_idx.set(idx);
            StartWorkout();
            goto(`/workout`);
          }}
          style="margin-right:5px; margin-left:5px;">{day.Name}</button
        >
      {/each}
    </details>
  {/if}
  <button class="breathe" on:click={() => (modalOpen = true)}
    >Start new workout plan</button
  >

  <button class="breathe" on:click={() => goto("/create_plan")}
    >Create new workout plan</button
  >
</div>

<style>
  .center {
    text-align: center;
    display: flex;
    flex-flow: column;
    justify-content: center;
  }

  .breathe {
    margin-top: 10px;
    margin-bottom: 10px;
  }
</style>
