<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { EventsOn } from "../lib/wailsjs/runtime";
  import { GetPlan } from "../lib/wailsjs/go/main/App";

  import PlanList from "./PlanList.svelte";

  let modalOpen = false;
  let plan_name;
  onMount(() => {
    GetPlan().then((res) => {
      plan_name = res;
    });
  });
</script>

<PlanList bind:modalOpen bind:plan_name />
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
