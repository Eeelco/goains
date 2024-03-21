<script>
  import ExercisesList from "./ExercisesList.svelte";
  import { GetExercises } from "../../lib/wailsjs/go/main/App";
  let plan_name = "";
  let filter_name = "";
  let nr_days = 3;
  let exercise_lists = new Array(nr_days).fill([]);
  let day_names = Array.from({ length: nr_days }, (_, i) => `Day ${i + 1}`);
  let modalOpen;

  function renameDay(i) {
    day_names[i] = prompt("New name", day_names[i]);
  }

  function addExercise(i) {
    modalOpen = true;
  }

  function removeExercise(i, exercise) {
    exercise_lists[i] = exercise_lists[i].filter((e) => e !== exercise);
  }

  document.addEventListener("keydown", (event) => {
    if (event.key === "Escape" && modalOpen) {
      modalOpen = false;
    }
  });
</script>

<ExercisesList bind:modalOpen />

<div role="group">
  <h1>New Plan</h1>
  <button
    on:click={() => console.log({ plan_name, exercise_lists, day_names })}
  >
    Save plan
  </button>
</div>
<input type="text" bind:value={plan_name} placeholder="Plan name" />
<button
  on:click={() => {
    nr_days += 1;
    exercise_lists[nr_days - 1] = [];
    day_names[nr_days - 1] = `Day ${nr_days}`;
  }}>Add day</button
>
<hr />
<div class="grid">
  {#each Array(nr_days) as _, i}
    <div>
      <details role="button" class="outline secondary" open={i == 0 || null}>
        <summary>{day_names[i]}</summary>
        <button on:click={() => renameDay(i)}>Rename day</button>
        <button on:click={() => addExercise(i)}>Add exercise</button>
        <div>
          <hr />
          <h3>Exercises</h3>
          {#each exercise_lists[i] as exercise}
            <article>
              <p>{exercise}</p>
              <button on:click={() => removeExercise(i, exercise)}
                >Remove</button
              >
            </article>
          {/each}
        </div>
      </details>
    </div>
  {/each}
</div>
