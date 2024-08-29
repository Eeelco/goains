<script>
  import { goto } from "$app/navigation";
  import { onDestroy } from "svelte";
  import { current_day_idx, plan, rest_timer } from "../stores.js";
  import { get } from "svelte/store";
  import { GetExercisesByIDs } from "$lib/wailsjs/go/main/App";
  import ExerciseCard from "../../lib/components/ExerciseCard.svelte";
  import TimeModal from "../../lib/components/TimeModal.svelte";

  let current_day;
  let current_exercise_idx = 0;
  let exercises = [];
  let is_break = false;
  let remaining_time = 0;

  plan.subscribe((p) => {
    current_day = p.Days[get(current_day_idx)];

    GetExercisesByIDs(
      current_day.ExerciseUnits.map((eu) => eu.ExerciseId)
    ).then((res) => {
      exercises = res;
    });
  });

  rest_timer.subscribe((rt) => {
    if (rt === 0) {
      is_break = false;
      return;
    }
    if (!is_break) {
      is_break = true;
    }
    remaining_time = rt;
    if (rt <= 0) {
      is_break = false;
    }
  });

  let exit_function = () => {
    if (confirm("Are you sure you want to exit?")) {
      goto("/");
    } else {
    }
  };

  let elapsed = 0;
  let elapsed_string = "";
  let last_time = window.performance.now();
  let frame;

  (function loop() {
    frame = requestAnimationFrame(loop);
    const now = window.performance.now();
    const delta = now - last_time;
    elapsed += delta;
    last_time = now;
    elapsed_string = new Date(elapsed).toISOString().substring(11, 19);

    if (is_break) {
      rest_timer.update((rt) => rt - delta/1000 );
    }
  })();

  onDestroy(() => {
    cancelAnimationFrame(frame);
  });
</script>

<div class="exercise-page">
  
<div role="group">
  <h2>{current_day.Name} <br /> {elapsed_string}</h2>
  <button class="contrast" on:click={exit_function}>Exit</button>
  <button>Save</button>
</div>
{#if exercises.length === 0}
  <p>No Exercises found</p>
{:else}
  <article>
    <header>
      <div class="center">
        <div class="breathe">
          <h4>{exercises[current_exercise_idx].Name}</h4>
        </div>
        <img
          src={`assets/img/${exercises[current_exercise_idx].Images[0]}`}
          alt={exercises[current_exercise_idx].Name}
        />

        <div role="group">
          {#if current_exercise_idx > 0}
            <button on:click={() => (current_exercise_idx -= 1)}
              >Previous</button
            >
          {/if}
          {#if current_exercise_idx < exercises.length - 1}
            <button on:click={() => (current_exercise_idx += 1)}>Next</button>
          {/if}
        </div>
      </div>
    </header>
    <body>
      <!-- <p>{exercises[current_exercise_idx].Instructions.join("\n")}</p> -->
      {#each current_day.ExerciseUnits[current_exercise_idx].Sets as set, i}
        <ExerciseCard
          {set}
          set_idx={i + 1}
          rest_time={current_day.ExerciseUnits[current_exercise_idx].Rest}
        />
      {/each}
    </body>
  </article>
{/if}
<TimeModal
  bind:modalOpen={is_break}
  bind:remaining_time
  max_time={current_day.ExerciseUnits[current_exercise_idx].Rest}
/>
</div>
<style>
  .exercise-page {
    position: relative;
  }
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
