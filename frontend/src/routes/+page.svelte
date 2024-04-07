<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { GetConfig } from "$lib/wailsjs/go/main/App";
  import Icon from "$lib/components/Icon.svelte";

  import PlanList from "./PlanList.svelte";
  import Settings from "./Settings.svelte";

  let modalOpen = false;
  let settingsOpen = false;
  let plan_name;
  let config = { DefaultNrReps: 0, DefaultNrSets: 0, DefaultRest: 0 };
  onMount(() => {
    GetConfig().then((res) => {
      config = res;
      plan_name = res.CurrentPlan;
    });
  });
</script>

<PlanList bind:modalOpen bind:plan_name />
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

  {#if plan_name !== ""}
    <button class="breathe" on:click={() => goto("/workout")}
      >Continue {plan_name}</button
    >
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
