<script>
  import { goto } from "$app/navigation";
  import { onDestroy, onMount } from "svelte";
  import { current_day_idx, plan, rest_timer } from "../stores.js";
  import { get } from "svelte/store";
  import {
    GetExercisesByIDs,
    WorkoutExitDialog,
    WorkoutSaveDialog,
    GetLastWorkout,
  } from "$lib/wailsjs/go/backend/App";
  import ExerciseCard from "$lib/components/ExerciseCard.svelte";
  import TimeModal from "$lib/components/TimeModal.svelte";

  let current_day;
  let current_exercise_idx = 0;
  let exercises = [];
  let is_break = false;

  plan.subscribe((p) => {
    current_day = p.Days[get(current_day_idx)];

    GetLastWorkout(current_day.Name).then((res) => {
      if (res.Data.Name === current_day.Name) {
        current_day = res.Data;
      }
    });

    for (let i = 0; i < current_day.ExerciseUnits.length; i++) {
      for (let j = 0; j < current_day.ExerciseUnits[i].Sets.length; j++) {
        current_day.ExerciseUnits[i].Sets[j].Done = false;
      }
    }

    GetExercisesByIDs(
      current_day.ExerciseUnits.map((eu) => eu.ExerciseId)
    ).then((res) => {
      exercises = res;
    });
  });

  // If rest timer is set to a positive value, it means we are in a break
  // If it is set to 0, it means the break is over
  rest_timer.subscribe((rt) => {
    if (rt <= 0) {
      is_break = false;
      return;
    }
    if (!is_break) {
      is_break = true;
    }
  });

  let stop_workout = () => {
    is_break = false;
    rest_timer.set(0);
    current_exercise_idx = 0;
    goto("/");
  };

  let exit_function = (save_progress) => {
    let future = save_progress ? WorkoutSaveDialog(current_day) : WorkoutExitDialog();
    future.then((res) => {
      if (res) {
        stop_workout();
      } else {
      }
    });
  };

  // Variables for showing workout timer
  let elapsed = 0;
  let elapsed_string = "";
  let last_time = window.performance.now();
  let frame;

  // Timer loop
  // Used to track length of workout and rest times
  (function loop() {
    frame = requestAnimationFrame(loop);
    const now = window.performance.now();
    const delta = now - last_time;
    elapsed += delta;
    last_time = now;
    elapsed_string = new Date(elapsed).toISOString().substring(11, 19);

    if (is_break) {
      rest_timer.update((rt) => rt - delta / 1000);
    }
  })();

  onDestroy(() => {
    cancelAnimationFrame(frame);
  });
</script>

<div class="exercise-page">
  <div role="group">
    <h2>{current_day.Name} <br /> {elapsed_string}</h2>
    <button class="contrast" on:click={() => {exit_function(false)}}>Exit</button>
    <button on:click={() => {exit_function(true)}}>Save</button>
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
          <!-- TODO: On clicking this image, it should open a modal with all images
             and instructions for the exercise -->
          <img
            src={`assets/img/${exercises[current_exercise_idx].Images[0]}`}
            alt={exercises[current_exercise_idx].Name}
          />

          <div role="group">
            {#if current_exercise_idx > 0}
              <button on:click={() => (current_exercise_idx -= 1)}
                >Show previous</button
              >
            {/if}
            {#if current_exercise_idx < exercises.length - 1}
              <button on:click={() => (current_exercise_idx += 1)}
                >Show next</button
              >
            {/if}
          </div>
        </div>
      </header>
      <body>
        {#each current_day.ExerciseUnits[current_exercise_idx].Sets as set, i}
          <!-- Shows number of reps and weight for each set -->
          <ExerciseCard
            bind:set
            set_idx={i + 1}
            rest_time={current_day.ExerciseUnits[current_exercise_idx].Rest}
          />
        {/each}
      </body>
    </article>
  {/if}
  <TimeModal
    bind:is_break
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
