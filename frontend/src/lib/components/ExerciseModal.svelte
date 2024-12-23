<script>
  import Carousel from "svelte-carousel";
  export let modal_open;
  export let modal_data;
  let images;
  $: {
    images = modal_data.Images.map((imgpath) => {
      return {
        alt: "Exercise execution",
        src: imgpath,
      };
    });
}
  let carousel;
</script>

<dialog open={modal_open}>
  {#if modal_data.Images.length > 0}
    <article>
      <header>
        <button
          aria-label="Close"
          rel="prev"
          on:click={() => {
            modal_data = {Images: []}
            modal_open = false;
          }}
        ></button>
        <p>
          <strong>{modal_data.Name}</strong>
        </p>
      </header>
      <Carousel bind:this={carousel} arrows={false} dots={false} autoplay={true}>
        {#each images as image}
          <div class="center">
            <img src={`assets/img/${image.src}`} alt={image.alt} />
          </div>
        {/each}
      </Carousel>
      <p style="text-align: justify;">{modal_data.Instructions.join('\n\n')}</p>
    </article>
  {/if}
</dialog>

<style>
    .center {
      text-align: center;
      display: flex;
      flex-flow: column;
      justify-content: center;
    }
  </style>