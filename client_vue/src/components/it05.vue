<template>
  <div class="queue-container">
    <div class="queue-header">IT 05-{{ page }}</div>
    <!-- IT 05-1: รับบัตรคิว -->
    <div v-if="page === 1" class="page">
      <button class="main-btn" @click="nextQueue">รับบัตรคิว</button>
      <button class="clear-btn" @click="gotoClear">ล้างคิว</button>
    </div>
    <!-- IT 05-2: แสดงหมายเลขคิว -->
    <div v-if="page === 2" class="page">
      <div style="font-size:22px;margin-bottom:20px;">หมายเลขคิว</div>
      <div class="queue-number">{{ queue }}</div>
      <div style="margin:20px 0;">วันที่ : {{ queueTime }}</div>
      <button class="back-btn" @click="page = 1">กลับไปหน้ารับบัตรคิว</button>
    </div>
    <!-- IT 05-3: ล้างคิว -->
    <div v-if="page === 3" class="page">
      <button class="clear-btn" @click="clearQueue">ล้างคิว</button>
      <div style="margin:20px 0;">หมายเลขคิวปัจจุบัน</div>
      <div class="queue-number">{{ queue }}</div>
      <button class="back-btn" @click="page = 1">กลับไปหน้ารับบัตรคิว</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'IT05',
  data() {
    return {
      page: 1,
      queue: '',
      queueTime: ''
    }
  },
  methods: {
    async nextQueue() {
      const res = await fetch('http://localhost:3000/api/it05/next-queue', { method: 'POST' });
      const data = await res.json();
      this.queue = data.current_queue || '';
      this.queueTime = new Date().toLocaleString('th-TH', { hour: '2-digit', minute: '2-digit', day: '2-digit', month: '2-digit', year: 'numeric' });
      this.page = 2;
    },
     async gotoClear() {
      // ดึงคิวปัจจุบันก่อนเข้าหน้าล้างคิว
      const res = await fetch('http://localhost:3000/api/it05/current-queue', { method: 'GET' });
      const data = await res.json();
      this.queue = data.CurrentQueue || 'A0';
      this.page = 3;
    },
    async clearQueue() {
      const res = await fetch('http://localhost:3000/api/it05/clear-queue', { method: 'POST' });
      const data = await res.json();
      this.queue = data.current_queue || 'A0';
    }
  }
}
</script>

<style scoped>
.queue-container {
  max-width: 500px;
  margin: 30px auto;
  border: 1px solid #222;
  border-radius: 6px;
  background: #fff;
  min-height: 400px;
  position: relative;
}
.queue-header {
  background: #27ae60;
  color: #fff;
  font-weight: bold;
  font-size: 18px;
  padding: 12px;
  text-align: center;
}
.page { 
    text-align: center; 
    margin-top: 60px; 
    display: flex;
    flex-direction: column;
    align-items: center;
}
.main-btn {
  font-size: 32px;
  padding: 40px 60px;
  background: #337ab7;
  color: #fff;
  border: none;
  border-radius: 12px;
  margin-bottom: 40px;
  cursor: pointer;
}
.clear-btn {
  font-size: 20px;
  padding: 16px 40px;
  background: #888;
  color: #fff;
  border: none;
  border-radius: 8px;
  margin-bottom: 20px;
  cursor: pointer;
}
.queue-number {
  font-size: 64px;
  font-weight: bold;
  margin: 20px 0;
}
.back-btn {
  font-size: 20px;
  padding: 16px 40px;
  background: #337ab7;
  color: #fff;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  margin-bottom: 20px;
}
</style>