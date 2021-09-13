<template>
<div class="camera-view">
  <h2>Camera test</h2>
  <img class="camera-image" v-if="imageObjectURL != ''" :src="imageObjectURL" />
  <button @click="swapInterval">{{ swapIntervalText }}</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      imageObjectURL: '',
      collect: false,
    }
  },
  methods: {
    fetchImage() {
      if(this.collect)
      {
        fetch(this.imageUrl)
          .then((response) => response.blob())
          .then((imageBlob) => {
            // Then create a local URL for that image and print it
            this.imageObjectURL = URL.createObjectURL(imageBlob)
            console.log(this.imageObjectURL)
            if(this.collect){
              setTimeout(this.fetchImage, 2000);
            }
          })
          .catch((e) => {
            console.log(e)
            if(this.collect){
              setTimeout(this.fetchImage, 2000);
            }
          })
      }
    },
    swapInterval() {
      if (!this.collect) {
        this.collect=true;
        this.fetchImage();
      } else {
        this.collect=false;
      }
    },
  },
  computed: {
    imageUrl() {
      return '/api/get-snapshot'
    },
    swapIntervalText() {
      if (!this.collect) {
        return 'Raccogli immagini'
      }
      return 'stop'
    },
  },
  beforeUnmount() {
    this.collect=false;
  },
}
</script>

<style>
.camera-image {
  max-width: 50%;
}
.camera-view {
  border-style: none solid solid solid;
  border-width: 1pt;
  padding-bottom: 10pt;
}
</style>
