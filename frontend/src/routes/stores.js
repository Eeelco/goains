import { writable } from "svelte/store";

export const config = writable({
  CurrentPlan: "",
  DefaultNrReps: 0,
  DefaultNrSets: 0,
  DefaultRest: 0,
});

export const plan = writable({ Days: [] });

export const current_day_idx = writable(0);

export const rest_timer = writable(0);
