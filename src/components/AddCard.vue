<template>
  <div>
    <select class="btnCard" id="cardListId" @change="updateCard()">
      <option value="0">Modify card ...</option>
      <option
        :key="card.id"
        v-for="card in cards"
        :value="JSON.stringify(card)"
      >
        {{ card.nameCard }}
      </option>
    </select>
  </div>
  <form @submit="onSubmit" class="add-form">
    <div class="form-control">
      <label>Name on card</label>
      <input
        type="text"
        v-model="card.nameCard"
        name="NameOnCard"
        placeholder="Name on card"
      />
    </div>
    <div class="form-control">
      <label>Card number</label>
      <input
        type="text"
        v-model="card.cardNumber"
        name="CardNumber"
        placeholder="Card number"
      />
    </div>
    <div class="form-control">
      <label>Expiration Date </label>
      <input
        type="text"
        v-model="card.expirationDate"
        name="ExpirationDate"
        placeholder="Expiration date (MM/YY)"
      />
    </div>
    <div class="form-control">
      <label>CVV</label>
      <input type="text" v-model="card.cvv" name="CVV" placeholder="CVV" />
    </div>
    <input
      v-if="!selected"
      type="submit"
      value="Save Card"
      class="btn btnSave"
    />
  </form>
  <div class="inner">
    <button v-if="selected" class="btn btnSave" @click.stop="onEdit">
      Save Edit Card
    </button>
    <button v-if="selected" class="btn btnCardDelete" @click.stop="cardDelete">
      Delete Card
    </button>
  </div>
</template>

<script>
import axios from "axios";
var IdCardCount = 1;
export default {
  name: "AddCard",
  props: {
    cards: Array,
  },
  data() {
    return {
      card: {
        id: "",
        nameCard: "",
        cardNumber: "",
        expirationDate: "",
        cvv: "",
      },
      selected: 0,
    };
  },
  methods: {
    updateCard() {
      let x = document.getElementById("cardListId").value;
      let cardObject = JSON.parse(x);
      console.log(x);
      this.card = { ...cardObject };
      this.selected = this.card.id;
    },

    //async
    onSubmit(e) {
      e.preventDefault();
      var cvvExpRegEx = /^([0-9]{3, 4})$/;
      var cardExpRegEx = /^((0[1-9]|1[0-2])\/[12]\d{3})$/;
      var visaRegEx = /^(?:4[0-9]{12}(?:[0-9]{3})?)$/;
      var mastercardRegEx = /^(?:5[1-5][0-9]{14})$/;
      var amexpRegEx = /^(?:3[47][0-9]{13})$/;
      var discovRegEx = /^(?:6(?:011|5[0-9][0-9])[0-9]{12})$/;

      if (!this.card.nameCard) {
        alert("Please add cardholder name");
        return;
      }
      if (!this.card.cardNumber) {
        alert("Please add card number");
        return;
      }
      if (
        visaRegEx.test(this.card.cardNumber) === false &&
        mastercardRegEx.test(this.card.cardNumber) === false &&
        amexpRegEx.test(this.card.cardNumber) === false &&
        discovRegEx.test(this.card.cardNumber) === false
      ) {
        alert("Please add a valid card number");
        return;
      }
      if (!this.card.expirationDate) {
        alert("Please add expiration date");
        return;
      }
      if (cardExpRegEx.test(this.card.expirationDate)) {
        alert("Invalid expiration date");
        return;
      }
      if (!this.card.cvv) {
        alert("Please add cvv");
        return;
      }
      if (cvvExpRegEx.test(this.card.cvv)) {
        alert("Invalid cvv");
        return;
      }
      this.$emit("add-card", {
        ...this.card,
        id: IdCardCount,
      });
      axios
        .post("http://localhost:9000/cardAdd", {
          id: IdCardCount,
          nameCard: this.card.nameCard,
          cardNumber: this.card.cardNumber,
          expirationDate: this.card.expirationDate,
          cvv: this.card.cvv,
        })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
      IdCardCount += 1;
      // let response = await axios.post("http://localhost:9000/card", {
      //   id: this.card.id,
      //   nameCard: this.card.nameCard,
      //   cardNumber: this.card.cardNumber,
      //   expirationDate: this.card.expirationDate,
      //   cvv: this.card.cvv,
      // });

      this.card = {
        nameCard: "",
        cardNumber: "",
        expirationDate: "",
        cvv: "",
      };
    },

    onEdit() {
      this.$emit("edit-card", this.card);
      const axios = require("axios").default;

      axios
        .post("http://localhost:9000/cardEdit", {
          id: this.card.id,
          nameCard: this.card.nameCard,
          cardNumber: this.card.cardNumber,
          expirationDate: this.card.expirationDate,
          cvv: this.card.cvv,
        })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
      this.card = {
        id: "",
        nameCard: "",
        cardNumber: "",
        expirationDate: "",
        cvv: "",
      };
      this.selected = 0;
      document.getElementById("cardListId").value = 0;
    },

    cardDelete() {
      this.$emit("delete-card", this.card);
      const axios = require("axios").default;

      axios
        .post("http://localhost:9000/cardDelete", {
          id: this.card.id,
        })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
      this.card = {
        id: "",
        nameCard: "",
        cardNumber: "",
        expirationDate: "",
        cvv: "",
      };
      this.selected = 0;
      document.getElementById("cardListId").value = 0;
    },

    async initCardId() {
      let response = await axios.get("http://localhost:9000/cardId");
      IdCardCount = response?.data;
      IdCardCount++;
      console.log(IdCardCount);
    },
  },
  async created() {
    this.initCardId();
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
  margin-top: 7px;
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

.btnCard {
  display: inline-block;
  background: #000;
  color: #fff;
  border: none;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  font-size: 15px;
  font-family: inherit;
  display: block;
  width: 310px;
  height: 40px;
  margin-left: 100px;
  margin-right: 390px;
  margin-bottom: 15px;
}
.btnCard:focus {
  outline: none;
}
.btnCard:active {
  transform: scale(0.98);
}
.btnCardDelete {
  width: 300px;
  height: 40px;
  display: inline-block;
  background: orangered;
  color: #fff;
  border: none;
  padding: 3px 7px;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  margin: 5px;
  margin-left: 210px;
  font-size: 15px;
  font-family: inherit;
}
.inner {
  display: inline-block;
}
</style>
