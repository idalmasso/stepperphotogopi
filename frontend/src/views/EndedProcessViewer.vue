<template>
  <div class="ended-process-viewer">
    <h2>Process {{ processName }}</h2>
    <vue-three-sixty 
      :amount=39
      :imagePath="imageDir"
      fileName="{index}.jpg"
    />
  </div>
</template>

<script>
export default {

  name: "EndedProcessViewer",
  props: {
    processName:String
  },
  computed: {
    imageDir() {
      //console.log(this.processName)
      return "/process-images/"+this.processName+"/";
    }
  },
  mounted() {

  }
}
</script>

<style>
@import '../assets/css/all.css';
</style>