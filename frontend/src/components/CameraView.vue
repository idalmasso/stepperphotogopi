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
      idCollect: -1
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
            
            if(this.collect){
              this.idCollect=setTimeout(this.fetchImage, 2000);
              this.$emit('setTimeout', this.idCollect);
            }
          })
          .catch(() => {
            if(this.collect){
              this.idCollect=setTimeout(this.fetchImage, 2000);
              this.$emit('setTimeout', this.idCollect);
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
        this.$emit('setTimeout', '');
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
  unmounted(){
    
    clearTimeout(this.idCollect)
  }
}
</script>

<style>
.camera-image {
  max-width: 50%;
}
.camera-view {
  border-style: solid;
  border-width: 1pt;
  padding-bottom: 10pt;
  margin: 1pt;
}
</style>
