<template>
  <div class="page">
    <div class="header">IT 09-1</div>

    <div class="post">
      <div class="user">
        <div class="avatar">C</div>
        <div>
          <strong>Change can</strong>
          <div class="time">16 October 2021 16:00</div>
        </div>
      </div>

      <img class="photo" src="/it09-pic.webp" />

      <div class="comment-box">
        <div class="avatar">B</div>
        <input
          v-model="input"
          placeholder="Comment"
          @keyup.enter="sendComment"
        />
      </div>
      <div v-for="c in comments" :key="c.id" class="comment">
        <div class="avatar">B</div>
        <div>
          <strong>{{ c.username }}</strong><br />
          {{ c.message }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"

const API = "http://localhost:3000/api/it09"
const input = ref("")
const comments = ref([])

onMounted(loadComments)
async function loadComments() {
  const res = await fetch(`${API}/comments`)
  comments.value = await res.json()
}

async function sendComment() {
  if (!input.value.trim()) return

  const res = await fetch(`${API}/insert-comment`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      message: input.value
    }),
  })

  if (!res.ok) {
    alert("ส่งไม่สำเร็จ")
    return
  }

  input.value = ""
  await loadComments()
}
</script>

<style scoped>
.page {
  width: 900px;
  margin: auto;
  border: 1px solid #333;
}
.header {
  background: #00b050;
  text-align: center;
  padding: 8px;
  font-weight: bold;
}
.post {
  padding: 20px;
}
.user {
  display: flex;
  gap: 10px;
}
.avatar {
  width: 40px;
  height: 40px;
  background: #2e6fd3;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}
.time {
  font-size: 12px;
  color: #666;
}
.photo {
  width: 100%;
  margin: 15px 0;
}
.comment-box {
  display: flex;
  gap: 10px;
  margin-bottom: 10px;
}
.comment-box input {
  flex: 1;
  padding: 6px;
}
.comment {
  display: flex;
  gap: 10px;
  margin-top: 8px;
}
</style>
