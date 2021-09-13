<template>
  <div class="test-page">
    <h1>Test page</h1>
    <MotorControl @setInterval="motorIntervalSet"/>
    <CameraView @setTimeout="imageTimeoutSet"/>
  </div>
</template>

<script>
// @ is an alias to /src
import MotorControl from '@/components/MotorControl.vue'
import CameraView from '@/components/CameraView.vue'

export default {
  name: 'TestView',
  data() {
    return {
      motorIntervalId:'',
      idCollectImage:''
    }
  },
  components: {
    MotorControl,
    CameraView,
  },
  methods:{
    motorIntervalSet(value){
      this.motorIntervalId=value;
    },
    imageTimeoutSet(value){
      this.idCollectImage=value
    }
  },
  beforeRouteLeave (to, from, next) {
    console.log("CameraView Clearing statusRequest")
      clearTimeout(this.idCollectImage);
      clearInterval(this.motorIntervalId);
      next();
  },
}
</script>
