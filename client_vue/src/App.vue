<template>
  <div>
    <div class="sidebar">
      <button v-for="n in 10" :key="n" :class="{ active: current === n }" @click="current = n">
        IT {{ n.toString().padStart(2, '0') }}
      </button>
    </div>
    <div class="main">
      <component :is="currentComponent" />
    </div>
  </div>
</template>

<script>
import IT01 from './components/it01.vue'
import IT02 from './components/it02.vue'
import IT03 from './components/it03.vue'
import IT04 from './components/it04.vue'
import IT05 from './components/it05.vue'
export default {
  components: { IT01, IT02, IT03, IT04, IT05 },
  data() {
    return {
      current: 1,
    }
  },
  computed: {
    currentComponent() {
      if (this.current === 1) return 'IT01'
      if (this.current === 2) return 'IT02'
      if (this.current === 3) return 'IT03'
      if (this.current === 4) return 'IT04'
      if (this.current === 5) return 'IT05'
      // ถ้ามี IT03 ถึง IT10 ให้เพิ่มตรงนี้
      return null
    },
  },
}
</script>

<style>
.sidebar {
  float: left;
  width: 120px;
  background: #eee;
  height: 100vh;
  padding: 10px 0;
}
.sidebar button {
  display: block;
  width: 100%;
  margin-bottom: 8px;
  padding: 8px;
  border: none;
  background: #ddd;
  cursor: pointer;
}
.sidebar button.active {
  background: #27ae60;
  color: #fff;
}
.main {
  margin-left: 130px;
  padding: 20px;
}
</style>
