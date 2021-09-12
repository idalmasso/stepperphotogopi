<template>
  <div class="home">
    <MotorControl />
    <CameraView />
  </div>
</template>

<script>
// @ is an alias to /src
import MotorControl from '@/components/MotorControl.vue'
import CameraView from '@/components/CameraView.vue'

export default {
  name: 'MotorView',
  components: {
    MotorControl,
    CameraView,
  },
}
</script>
