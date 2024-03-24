<script>
  import ExercisesList from "./ExercisesList.svelte";
  import { GetExercises, SavePlan } from "../../lib/wailsjs/go/main/App";
  let plan_name = "";
  let filter_name = "";
  let nr_days = 3;
  let exercise_lists = [...new Array(nr_days)].map(() => []);
  let day_names = Array.from({ length: nr_days }, (_, i) => `Day ${i + 1}`);
  let modalOpen;
  let current_day = 0;

  function renameDay(i) {
    res = prompt("New name", day_names[i]);
    if (res) {
      day_names[i] = res;
    }
  }

  function addExercise(i) {
    current_day = i;
    modalOpen = true;
  }

  function removeExercise(i, exercise) {
    exercise_lists[i] = exercise_lists[i].filter((e) => e !== exercise);
  }

  function savePlan() {
    const plan = {
      Name: plan_name ? plan_name : "New plan",
      Days: exercise_lists.map((exercises, i) => ({
        Name: day_names[i],
        ExerciseUnits: exercises,
      })),
    };
    SavePlan(plan).then(() => {
      alert("Plan saved successfully");
    });
  }

  document.addEventListener("keydown", (event) => {
    if (event.key === "Escape" && modalOpen) {
      modalOpen = false;
    }
  });
</script>

<ExercisesList bind:modalOpen bind:exercise_lists bind:current_day />

<div role="group">
  <h1>New Plan</h1>
  <button on:click={savePlan}> Save plan </button>
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
            <details>
              <summary>{exercise.ExerciseName}</summary>
              <label>Rest</label>
              <input type="number" bind:value={exercise.Rest} min="1" />
              {#each Array(exercise.Sets) as _, j}
                <label>Set {j + 1} Reps</label>
                <input
                  type="number"
                  bind:value={exercise.Sets[j].Repetitions}
                  min="1"
                />
                <label>Set {j + 1} Weight</label>
                <input
                  type="number"
                  bind:value={exercise.Sets[j].Weight}
                  min="0"
                  step="0.1"
                />
              {/each}
              <button on:click={() => removeExercise(i, exercise)}
                >Remove</button
              >
            </details>
          {/each}
        </div>
      </details>
    </div>
  {/each}
</div>
