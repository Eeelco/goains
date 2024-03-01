<script>
  import { goto } from "$app/navigation";
  import { EventsOn } from "../../lib/wailsjs/runtime";

  let i = 0;
  let max = 1;

  EventsOn("download_db_progress", (progress) => {
    i = progress[0];
    max = progress[1];
  });
  EventsOn("download_db_done", () => {
    goto("/");
  });
</script>

<label for="files">Downloading files...</label>
<progress id="files" value={i} {max}></progress>
