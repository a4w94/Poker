//open load
const MaxRound = 79
var Loading = [];

//input
$("body").on("click", "#button", function () {
  InputRound = $("[name='round']:checked").val()

  var loadRound = true

  let checked = document.querySelector('[name=round]:checked')
  if (!checked) {
    alert("局數不得為空")
    loadRound = false
  }
  var data = {
    dataAmount: $("[name='dataAmount']").val(),
    round: $("[name='round']:checked").val(),
    Banker_and_Tie_Win: $("#Banker_and_Tie_Win").val(),
    Banker_Win: $("#Banker_Win").val(),
    Player_Win: $("#Player_Win").val(),
    Tie_Win: $("#Tie_Win").val(),
    Banker_Pair: $("#Banker_Pair").val(),
    Player_Pair: $("#Player_Pair").val(),
    Banker_28gang: $("#Banker_28gang").val(),
    Player_28gang: $("#Player_28gang").val(),
    Banker_10point: $("#Banker_10point").val(),
    Player_10point: $("#Player_10point").val(),
    White_1: $("#White_1").val(),
    White_2: $("#White_2").val(),
    White_3: $("#White_3").val(),
    White_4: $("#White_4").val(),

  }
  console.log(data)

  Loading = document.getElementById('loading')

  //thmrtp
  var oReq = new XMLHttpRequest();
  oReq.addEventListener("load", GetTheorRTP);
  oReq.open("POST", "/api/v1/CalThmRTP");
  oReq.send(JSON.stringify(data));

  //range rate
  if (loadRound == true) {

    Loading.innerHTML = `計算中...`
    var oReq1 = new XMLHttpRequest();
    oReq1.addEventListener("load", GetRangeRate);
    oReq1.open("POST", "/api/v1/CalRangeRate");
    oReq1.send(JSON.stringify(data));
  }


});



//Start Load Theory Rate
var TheoryRate = []
var TheoryRTP = []
var RangeRate = []
var InputRound = []

function GetRangeRate() {


  var tmp = JSON.parse(this.responseText);
  console.log('rate', tmp)
  RangeRate = tmp

  var table_of_rate = document.getElementById('RoundResult')

  console.log("result", table_of_rate)

  var output_decimal = 4


  table_of_rate.innerHTML = ``

  for (var i = 1; i > 0; i++) {
    var row = `
          <tr>
           <td>第${InputRound * i}局</td>
           <td >${NumberToString(tmp[InputRound * i].Banker_and_Tie_Win, output_decimal)}</td>
           <td>${NumberToString(tmp[InputRound * i].Banker_Win, output_decimal)}</td>
           <td>${NumberToString(tmp[InputRound * i].Player_Win, output_decimal)}</td>
           <td>${NumberToString(tmp[InputRound * i].Tie_Win, output_decimal)}</td>
           <td>${NumberToString(tmp[InputRound * i].Banker_Pair, output_decimal)}</td>
           <td>${NumberToString(tmp[InputRound * i].Player_Pair, output_decimal)}</td>
           <td>${NumberToString(tmp[InputRound * i].Banker_28gang, output_decimal)}</td>
           <td>${NumberToString(tmp[InputRound * i].Player_28gang, output_decimal)}</td>
           <td>${NumberToString(tmp[InputRound * i].Banker_10point, output_decimal)}</td>
           <td>${NumberToString(tmp[InputRound * i].Player_10point, output_decimal)}</td>
           <td colspan="4">${NumberToString(tmp[InputRound * i].WhiteBetArea, output_decimal)}</td>
           </tr>
           `
    table_of_rate.innerHTML += row

    if (InputRound * (i + 1) > MaxRound) {
      break
    }
  }

  Loading.innerHTML = `計算完成`
}


///
function GetTheorRTP() {
  var tmp = JSON.parse(this.responseText);
  console.log('RTP', tmp)
  TheoryRTP = tmp


  var table_of_theory_rate = document.getElementById('thmRTP')

  var output_decimal = 3

  var row = `
           <td>理論RTP</td>
           <td >${NumberToString(tmp.Banker_and_Tie_Win, output_decimal)}</td>
           <td>${NumberToString(tmp.Banker_Win, output_decimal)}</td>
           <td>${NumberToString(tmp.Player_Win, output_decimal)}</td>
           <td>${NumberToString(tmp.Tie_Win, output_decimal)}</td>
           <td>${NumberToString(tmp.Banker_Pair, output_decimal)}</td>
           <td>${NumberToString(tmp.Player_Pair, output_decimal)}</td>
           <td>${NumberToString(tmp.Banker_28gang, output_decimal)}</td>
           <td>${NumberToString(tmp.Player_28gang, output_decimal)}</td>
           <td>${NumberToString(tmp.Banker_10point, output_decimal)}</td>
           <td>${NumberToString(tmp.Player_10point, output_decimal)}</td>
           <td colspan="4">${NumberToString(tmp.WhiteBetArea, output_decimal)}</td>
          `
  table_of_theory_rate.innerHTML = row

}
function GetTheorRate() {
  var tmp = JSON.parse(this.responseText);


  TheoryRate = tmp


  var table_of_theory_rate = document.getElementById('thmRate')

  var output_decimal = 3

  var row = `</tr>
           <td>理論機率</td>
           <td >${NumberToString(tmp.Banker_and_Tie_Win, output_decimal)}</td>
           <td>${NumberToString(tmp.Banker_Win, output_decimal)}</td>
           <td>${NumberToString(tmp.Player_Win, output_decimal)}</td>
           <td>${NumberToString(tmp.Tie_Win, output_decimal)}</td>
           <td>${NumberToString(tmp.Banker_Pair, output_decimal)}</td>
           <td>${NumberToString(tmp.Player_Pair, output_decimal)}</td>
           <td>${NumberToString(tmp.Banker_28gang, output_decimal)}</td>
           <td>${NumberToString(tmp.Player_28gang, output_decimal)}</td>
           <td>${NumberToString(tmp.Banker_10point, output_decimal)}</td>
           <td>${NumberToString(tmp.Player_10point, output_decimal)}</td>
           <td>${NumberToString(tmp.White_1, output_decimal)}</td>
           <td>${NumberToString(tmp.White_2, output_decimal)}</td>
           <td>${NumberToString(tmp.White_3, output_decimal)}</td>
           <td>${NumberToString(tmp.White_4, output_decimal)}</td>
      <tr>`
  table_of_theory_rate.innerHTML += row

}



function NumberToString(number, decimal) {
  var result = (Math.round(number * Math.pow(10, decimal) * 100) / Math.pow(10, decimal)) + '%'

  return result
}

var oReq = new XMLHttpRequest();
oReq.addEventListener("load", GetTheorRate);
oReq.open("GET", "/api/v1/ThmRate");
oReq.send();





