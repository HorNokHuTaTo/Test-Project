<template>
  <div>
    <div class="header">IT 03-1</div>
    <div class="actions">
      <button class="approve-btn" @click="openApproveModal" :disabled="!canApprove">อนุมัติ</button>
      <button class="reject-btn" @click="openRejectModal" :disabled="!canApprove">ไม่อนุมัติ</button>
    </div>
    <table class="doc-table">
      <thead>
        <tr>
          <th><input type="checkbox" @change="toggleAll" :checked="allChecked" /></th>
          <th>รายการ</th>
          <th>เหตุผล</th>
          <th>สถานะเอกสาร</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="doc in docs" :key="doc.id">
          <td>
            <input
              type="checkbox"
              :disabled="doc.status !== 'รออนุมัติ'"
              v-model="doc.checked"
            />
          </td>
          <td>{{ doc.name }}</td>
          <td>{{ doc.reason }}</td>
          <td>{{ doc.status }}</td>
        </tr>
      </tbody>
    </table>
    <div v-if="showApproveModal" class="modal-overlay">
      <div class="modal">
        <div class="modal-header">ยืนยันการอนุมัติ</div>
        <div class="modal-body">
          <label>เหตุผล :</label>
          <textarea v-model="modalReason"></textarea>
          <div v-if="reasonError" class="error-msg">{{ reasonError }}</div>
        </div>
        <div class="modal-actions">
          <button class="approve-btn" @click="approveSelected">อนุมัติ</button>
          <button class="cancel-btn" @click="closeModal">ยกเลิก</button>
        </div>
      </div>
    </div>
    <div v-if="showRejectModal" class="modal-overlay">
      <div class="modal">
        <div class="modal-header">ยืนยันการไม่อนุมัติ</div>
        <div class="modal-body">
          <label>เหตุผล :</label>
          <textarea v-model="modalReason"></textarea>
          <div v-if="reasonError" class="error-msg">{{ reasonError }}</div>
        </div>
        <div class="modal-actions">
          <button class="reject-btn" @click="rejectSelected">ไม่อนุมัติ</button>
          <button class="cancel-btn" @click="closeModal">ยกเลิก</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
async function updateStatus(ids, status, description) {
  const res = await fetch('https://test-project-0w71.onrender.com/api/it03/update-status', {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      ids,
      status,
      description
    })
  });
  
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || 'Update failed');
  }
  return await res.json();
}

async function fetchDocs() {
  const res = await fetch('https://test-project-0w71.onrender.com/api/it03/documents');
  if (!res.ok) throw new Error('โหลดข้อมูลไม่สำเร็จ');
  return await res.json();
}

export default {
  data() {
    return {
      docs: [],
      showApproveModal: false,
      showRejectModal: false,
      modalReason: "",
      errorMsg: "",
        reasonError: ""
    };
  },
  computed: {
    canApprove() {
      return this.docs.some(doc => doc.checked && doc.status === "รออนุมัติ");
    },
    allChecked() {
      return this.docs.filter(doc => doc.status === "รออนุมัติ").every(doc => doc.checked);
    }
  },
  methods: {
    async loadDocs() {
      try {
        const data = await fetchDocs();
        this.docs = data.map(doc => ({
  id: doc.id,
  name: doc.title,         
  reason: doc.description, 
  status: doc.status,
  checked: false
}));
      } catch (e) {
        this.errorMsg = e.message;
      }
    },
    toggleAll(e) {
      const checked = e.target.checked;
      this.docs.forEach(doc => {
        if (doc.status === "รออนุมัติ") doc.checked = checked;
      });
    },
    openApproveModal() {
      if (this.canApprove) {
        this.modalReason = "";
        this.showApproveModal = true;
      }
    },
    openRejectModal() {
      if (this.canApprove) {
        this.modalReason = "";
        this.reasonError = "";
        this.showRejectModal = true;
      }
    },
    closeModal() {
      this.showApproveModal = false;
      this.showRejectModal = false;
    },
    async approveSelected() {
      const ids = this.docs
        .filter(doc => doc.checked && doc.status === "รออนุมัติ")
        .map(doc => doc.id)
      if (ids.length === 0) 
      return;
    if (!this.modalReason.trim()) {
    this.reasonError = "กรุณากรอกเหตุผล";
    return;
  }
      try {
        await updateStatus(ids, "อนุมัติ", this.modalReason);
        await this.loadDocs();
        this.closeModal();
      } catch (e) {
        alert(e.message);
      }
    },
    async rejectSelected() {
  const ids = this.docs
    .filter(doc => doc.checked && doc.status === "รออนุมัติ")
    .map(doc => doc.id)
    .filter(id => !!id); 
  if (ids.length === 0) 
  return;
  if (!this.modalReason.trim()) {
    this.reasonError = "กรุณากรอกเหตุผล";
    return;
  }
  try {
    await updateStatus(ids, "ไม่อนุมัติ", this.modalReason);
    await this.loadDocs();
    this.closeModal();
  } catch (e) {
    alert(e.message);
  }
}
  },
  mounted() {
    this.loadDocs();
  }
};
</script>


<style scoped>
.header {
  background: #27ae60;
  color: #fff;
  text-align: center;
  padding: 10px;
  font-weight: bold;
  border-radius: 6px 6px 0 0;
  margin-bottom: 10px;
}
.actions {
  margin-bottom: 10px;
  text-align: left;
}
.approve-btn {
  background: #27ae60;
  color: #fff;
  border: none;
  padding: 6px 16px;
  border-radius: 4px;
  margin-right: 8px;
  font-weight: bold;
  cursor: pointer;
}
.reject-btn {
  background: #e74c3c;
  color: #fff;
  border: none;
  padding: 6px 16px;
  border-radius: 4px;
  font-weight: bold;
  cursor: pointer;
}
.cancel-btn {
  background: #bbb;
  color: #222;
  border: none;
  padding: 6px 16px;
  border-radius: 4px;
  font-weight: bold;
  cursor: pointer;
}
.doc-table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
}
.doc-table th, .doc-table td {
  border: 1px solid #222;
  padding: 6px 10px;
  text-align: center;
}
.doc-table th {
  background: #2980b9;
  color: #fff;
  font-weight: bold;
}
.modal-overlay {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}
.modal {
  background: #fff;
  border-radius: 8px;
  min-width: 350px;
  min-height: 180px;
  box-shadow: 0 2px 8px #222;
  padding: 20px;
}
.modal-header {
  background: #2980b9;
  color: #fff;
  padding: 8px;
  font-weight: bold;
  border-radius: 6px 6px 0 0;
  margin-bottom: 10px;
  text-align: center;
}
.modal-body {
  margin-bottom: 16px;
}
.modal-body label {
  display: block;
  margin-bottom: 6px;
}
.modal-body textarea {
  width: 100%;
  min-height: 60px;
  border-radius: 4px;
  border: 1px solid #bbb;
  padding: 6px;
}
.modal-actions {
  text-align: right;
}
.error-msg {
  color: red;
  margin-top: 4px;
  text-align: left;
}
</style>

