<template>
  <form @submit="onSubmit" class="add-form">
    <div class="form-control">
      <label>Name on card</label>
      <input
        type="text"
        v-model="nameCard"
        name="NameOnCard"
        placeholder="Name on card"
      />
    </div>
    <div class="form-control">
      <label>Card number</label>
      <input
        type="text"
        v-model="cardNumber"
        name="CardNumber"
        placeholder="Card number"
      />
    </div>
    <div class="form-control">
      <label>Expiration Date </label>
      <input
        type="text"
        v-model="expirationDate"
        name="ExpirationDate"
        placeholder="Expiration date (MM/YYYY)"
      />
    </div>
    <div class="form-control">
      <label>CVV</label>
      <input type="text" v-model="cvv" name="CVV" placeholder="CVV" />
    </div>
    <input type="submit" value="Save Card" class="btn btnSave" />
  </form>
</template>

<script>
export default {
  name: "AddCard",
  data() {
    return {
      nameCard: "",
      cardNumber: "",
      expirationDate: "",
      cvv: "",
    };
  },
  methods: {
    onSubmit(e) {
      e.preventDefault();
      var cvvExpRegEx = /^([0-9]{3, 4})$/;
      var cardExpRegEx = /^((0[1-9]|1[0-2])\/[12]\d{3})$/;
      var visaRegEx = /^(?:4[0-9]{12}(?:[0-9]{3})?)$/;
      var mastercardRegEx = /^(?:5[1-5][0-9]{14})$/;
      var amexpRegEx = /^(?:3[47][0-9]{13})$/;
      var discovRegEx = /^(?:6(?:011|5[0-9][0-9])[0-9]{12})$/;

      if (!this.nameCard) {
        alert("Please add cardholder name");
        return;
      }
      if (!this.cardNumber) {
        alert("Please add card number");
        return;
      }
      if (
        visaRegEx.test(this.cardNumber) === false &&
        mastercardRegEx.test(this.cardNumber) === false &&
        amexpRegEx.test(this.cardNumber) === false &&
        discovRegEx.test(this.cardNumber) === false
      ) {
        alert("Please add a valid card number");
        return;
      }
      if (!this.expirationDate) {
        alert("Please add expiration date");
        return;
      }
      if (cardExpRegEx.test(this.expirationDate)) {
        alert("Invalid expiration date");
        return;
      }
      if (!this.cvv) {
        alert("Please add cvv");
        return;
      }
      if (cvvExpRegEx.test(this.cvv)) {
        alert("Invalid cvv");
        return;
      }
      const newCard = {
        // id: Math.floor(Math.random() * 100000),
        nameCard: this.nameCard,
        cardNumber: this.cardNumber,
        expirationDate: this.expirationDate,
        cvv: this.cvv,
      };
      this.$emit("add-card", newCard);
      this.nameCard = "";
      this.cardNumber = "";
      this.expirationDate = "";
      this.cvv = "";
    },
  },
};
</script>

<style scoped>
.add-form {
  margin-bottom: 40px;
}
.form-control {
  float: left;
  margin: 20px 100px;
  margin-top: 0px;
}
.form-control label {
  display: block;
  color: cadetblue;
}
.form-control input {
  width: 300px;
  height: 35px;
  margin: 5px;
  padding: 3px 7px;
  font-size: 17px;
}
.form-control-check label {
  flex: 1;
  color: cadetblue;
}
.form-control-check input {
  flex: 2;
  height: 20px;
}
.btnSave {
  width: 300px;
  height: 40px;
  margin: 5px;
  padding: 3px 7px;
  margin-left: 100px;
}
</style>
