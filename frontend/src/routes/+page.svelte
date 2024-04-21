<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { GetConfig, GetPlan } from "$lib/wailsjs/go/main/App";
  import Icon from "$lib/components/Icon.svelte";

  import PlanList from "./PlanList.svelte";
  import Settings from "./Settings.svelte";

  const rotateArray = (arr, k) => arr.slice(k).concat(arr.slice(0, k));

  let modalOpen = false;
  let settingsOpen = false;
  let config = {
    CurrentPlan: "",
    DefaultNrReps: 0,
    DefaultNrSets: 0,
    DefaultRest: 0,
  };
  let plan = { Days: [] };
  let days = [];
  onMount(() => {
    GetConfig().then((res) => {
      config = res;
      GetPlan(config.CurrentPlan).then((res) => {
        plan = res;
        days = rotateArray(plan.Days, config.NextDayIdx);
      });
    });
  });
</script>

<PlanList bind:modalOpen bind:config />
<Settings bind:settingsOpen bind:config />
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
  <h1>Goainz</h1>

  {#if config.CurrentPlan !== ""}
    <details open>
      <summary role="button" class="secondary"
        >Continue {config.CurrentPlan}</summary
      >
      {#each days as day}
        <!-- <p> -->
        <button class="breathe" on:click={() => goto(`/workout/${day.Name}`)}
          >{day.Name}</button
        >
        <!-- </p> -->
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
