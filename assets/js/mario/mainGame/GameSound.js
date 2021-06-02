function GameSound() {
  var coin;
  var powerUpAppear;
  var powerUp;
  var marioDie;
  var killEnemy;
  var stageClear;
  var bullet;
  var powerDown;
  var jump;

  var that = this;

  this.init = function() {
    coin = new Audio('../../../assets/sounds/coin.wav');
    powerUpAppear = new Audio('../../../assets/sounds/power-up-appear.wav');
    powerUp = new Audio('../../../assets/sounds/power-up.wav');
    marioDie = new Audio('../../../assets/sounds/mario-die.wav');
    killEnemy = new Audio('../../../assets/sounds/kill-enemy.wav');
    stageClear = new Audio('../../../assets/sounds/stage-clear.wav');
    bullet = new Audio('../../../assets/sounds/bullet.wav');
    powerDown = new Audio('../../../assets/sounds/power-down.wav');
    jump = new Audio('../../../assets/sounds/jump.wav');
  };

  this.play = function(element) {
    if (element == 'coin') {
      coin.pause();
      coin.currentTime = 0;
      coin.play();
    } else if (element == 'powerUpAppear') {
      powerUpAppear.pause();
      powerUpAppear.currentTime = 0;
      powerUpAppear.play();
    } else if (element == 'powerUp') {
      powerUp.pause();
      powerUp.currentTime = 0;
      powerUp.play();
    } else if (element == 'marioDie') {
      marioDie.pause();
      marioDie.currentTime = 0;
      marioDie.play();
    } else if (element == 'killEnemy') {
      killEnemy.pause();
      killEnemy.currentTime = 0;
      killEnemy.play();
    } else if (element == 'stageClear') {
      stageClear.pause();
      stageClear.currentTime = 0;
      stageClear.play();
    } else if (element == 'bullet') {
      bullet.pause();
      bullet.currentTime = 0;
      bullet.play();
    } else if (element == 'powerDown') {
      powerDown.pause();
      powerDown.currentTime = 0;
      powerDown.play();
    } else if (element == 'jump') {
      jump.pause();
      jump.currentTime = 0;
      jump.play();
    }
  };
}
