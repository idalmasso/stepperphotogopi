<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">
              <h2>Select item</h2>
            </slot>
          </div>
          <div class="modal-body">
            <slot name="body">
              <table>
                <tr>
                  <td>
                    <label for="item">articolo</label>
                  </td>
                  <td>
                    <select id="item" v-model="item">
                      <option v-for="i in items" :key="i.id" :value="i.id">{{
                        i.name
                      }}</option>
                    </select>
                  </td>
                </tr>
              </table>
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              <button class="modal-default-button" @click="cancelClicked">
                CANCEL
              </button>
              <button class="modal-default-button" @click="okClicked">
                OK
              </button>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: 'ItemSelectorDialog',
  props: { itemsUrl: String },
  data() {
    return {
      item: '',
      items: [],
    }
  },
  methods: {
    okClicked() {
      this.$emit('selected', this.item)
    },
    cancelClicked() {
      this.$emit('close')
    },
  },
  mounted() {
    fetch(this.itemsUrl)
      .then((data) => {
        return data.json()
      })
      .then((jsonData) => {
        this.items = jsonData
      })
  },
}
</script>

<style scoped>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 300px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  margin: 20px 0;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>
