<template>
  <div class="ended-process-viewer">
    <h2>Processo {{ processName }}</h2>
    <vue-three-sixty
      :amount="numImages"
      :imagePath="imageDir"
      fileName="{index}.webp"
    />
  </div>
</template>

<script>
export default {
  name: 'EndedProcessViewer',
  props: {
    processName: String,
    numImages: Number,
  },
  computed: {
    imageDir() {
      //console.log(this.processName)
      return '/process-images/' + this.processName + '/'
    },
  },
}
</script>

<style>
@import '../assets/css/all.css';
</style>
