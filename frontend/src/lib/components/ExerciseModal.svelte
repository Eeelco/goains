<script>
    export let modal_open = true
    export let modal_data = [];
    let images = modal_data.Images.map((imgpath) => {
        return {
            alt: 'Exercise execution',
            src: imgpath, 
        };
    })
</script>

<dialog open={modal_open}>
    {#if modal_data.length > 0}
    <article>
        <header>
            <button aria-label="Close" rel="prev" on:click={() => {modal_open = false;}}></button>
            <p>
                <strong>{modal_data.Name}</strong>
            </p>
        </header>
    </article>
    {/if}
</dialog>