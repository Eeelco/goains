<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { EventsOn } from "../lib/wailsjs/runtime";
  import { HasPlan } from "../lib/wailsjs/go/main/App";

  import PlanList from "./PlanList.svelte";

  let modalOpen = false;
  let has_plan;
  onMount(() => {
    HasPlan().then((res) => {
      has_plan = res;
    });
  });
</script>

<PlanList bind:modalOpen />
<div class="center">
  <h1>Goainz</h1>

  {#if has_plan === true}
    <button class="breathe" on:click={() => goto("/workout")}
      >Continue workout</button
    >
  {:else}
    <button class="breathe" on:click={() => (modalOpen = true)}
      >Start new workout plan</button
    >
  {/if}

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
