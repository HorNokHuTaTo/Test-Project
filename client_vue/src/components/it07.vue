<template>
  <div class="page">
    <div class="header">IT 07-1</div>

    <div class="form-row">
      <label>รหัสสินค้า</label>
      <input
        v-model="skuInput"
        placeholder="XXXXX-XXXXX-XXXXX-XXXXX-XXXXX-XXXXX"
      />
      <button class="btn-add" @click="addSku">ADD</button>
    </div>

    <table class="sku-table">
      <thead>
        <tr>
          <th>Id</th>
          <th>รหัสสินค้า (30 หลัก)</th>
          <th>View</th>
          <th>Delete</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(sku, index) in skus" :key="sku.id">
          <td>{{ index + 1 }}</td>
          <td>{{ sku.sku_code }}</td>
          <td>
            <button class="btn-view" @click="openQr(sku)">QR</button>
          </td>
          <td>
            <button class="btn-del" @click="openDelete(index)">ลบ</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="showQrModal" class="modal-backdrop">
      <div class="modal">
        <img
          :src="`data:image/png;base64,${qrTarget?.qr_code}`"
          class="qr-big"
        />
        <div class="modal-actions">
          <button class="btn-cancel" @click="showQrModal = false">Close</button>
        </div>
      </div>
    </div>
    <div v-if="showDeleteModal" class="modal-backdrop">
      <div class="modal">
        <p>
          ต้องการลบข้อมูล รหัสสินค้า
          <strong>{{ deleteTarget?.sku_code }}</strong> หรือไม่ ?
        </p>
        <div class="modal-actions">
          <button class="btn-ok" @click="confirmDelete">ตกลง</button>
          <button class="btn-cancel" @click="showDeleteModal = false">ยกเลิก</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"

const skuInput = ref("")
const skus = ref([])

const showQrModal = ref(false)
const showDeleteModal = ref(false)

const qrTarget = ref(null)
const deleteTarget = ref(null)

const API = "https://test-project-0w71.onrender.com/api/it07" 

onMounted(loadSkus)

async function loadSkus() {
  const res = await fetch(`${API}/skus`)
  skus.value = await res.json()
}

async function addSku() {
  if (!skuInput.value) return

  const res = await fetch(`${API}/insert-sku`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ sku_code: skuInput.value }),
  })

  if (!res.ok) {
    if ( res.status === 400 ) {
      const data = await res.json()
      alert(data.error)
      return
    }

    alert("Insert failed")
    return
  }

  skuInput.value = ""
  await loadSkus()
}

function openQr(sku) {
  qrTarget.value = sku
  showQrModal.value = true
}

function openDelete(index) {
  deleteTarget.value = skus.value[index]
  showDeleteModal.value = true
}

async function confirmDelete() {
  const id = deleteTarget.value.id

  const res = await fetch(`${API}/delete-sku?id=${encodeURIComponent(id)}`, {
    method: "DELETE",
  })

  if (!res.ok) {
    alert("Delete failed")
    return
  }

  showDeleteModal.value = false
  await loadSkus()
}
</script>

<style scoped>
.page {
  width: 900px;
  margin: 20px auto;
  border: 1px solid #333;
}

.header {
  background: #00b050;
  text-align: center;
  padding: 8px;
  font-weight: bold;
}

.form-row {
  padding: 10px;
  display: flex;
  gap: 10px;
  align-items: center;
}

.form-row input {
  flex: 1;
  padding: 6px;
}

.btn-add {
  background: #2e6fd3;
  color: #fff;
  border: none;
  padding: 6px 14px;
  cursor: pointer;
}

.sku-table {
  width: 100%;
  border-collapse: collapse;
}

.sku-table th {
  background: #00b0f0;
  border: 1px solid #333;
  padding: 6px;
}

.sku-table td {
  border: 1px solid #333;
  padding: 6px;
  text-align: center;
}

.btn-view {
  background: #2ecc71;
  color: white;
  border: none;
  padding: 4px 10px;
  border-radius: 4px;
  cursor: pointer;
}

.btn-del {
  background: red;
  color: white;
  border: none;
  padding: 4px 10px;
  border-radius: 4px;
  cursor: pointer;
}

.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.3);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal {
  background: #fff;
  padding: 20px 30px;
  border-radius: 6px;
  text-align: center;
}

.qr-big {
  width: 260px;
  height: 260px;
}

.modal-actions {
  margin-top: 15px;
  display: flex;
  justify-content: center;
  gap: 10px;
}

.btn-ok {
  background: #2e6fd3;
  color: white;
  border: none;
  padding: 6px 16px;
  cursor: pointer;
}

.btn-cancel {
  background: red;
  color: white;
  border: none;
  padding: 6px 16px;
  cursor: pointer;
}
</style>
