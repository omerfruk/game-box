function Preloader() {
  var view = View.getInstance();

  var loadingPercentage;

  var imageSources;
  var soundSources;

  var that = this;

  this.init = function() {
    loadingPercentage = view.create('div');

    view.addClass(loadingPercentage, 'loading-percentage');
    view.setHTML(loadingPercentage, '0%');
    view.appendToBody(loadingPercentage);

    imageSources = {
      1: '../img/mario/back-btn.png',
      2: '../img/mario/bg.png',
      3: '../img/mario/bullet.png',
      4: '../img/mario/clear-map-btn.png',
      5: '../img/mario/coin.png',
      6: '../img/mario/delete-all-btn.png',
      7: '../img/mario/editor-btn.png',
      8: '../img/mario/elements.png',
      9: '../img/mario/enemies.png',
      10: '../img/mario/flag-pole.png',
      11: '../img/mario/flag.png',
      12: '../img/mario/start-screen.png',
      13: '../img/mario/grid-large-btn.png',
      14: '../img/mario/grid-medium-btn.png',
      15: '../img/mario/grid-small-btn.png',
      16: '../img/mario/grid.png',
      17: '../img/mario/lvl-size.png',
      18: '../img/mario/mario-head.png',
      19: '../img/mario/mario-sprites.png',
      20: '../img/mario/powerups.png',
      21: '../img/mario/save-map-btn.png',
      22: '../img/mario/saved-btn.png',
      23: '../img/mario/slider-left.png',
      24: '../img/mario/slider-right.png',
      25: '../img/mario/start-btn.png'
      
    };
    that.loadImages(imageSources);
  };

  this.loadImages = function(imageSources) {
    var images = {};
    var loadedImages = 0;
    var totalImages = 0;

    for (var key in imageSources) {
      totalImages++;
    }

    for (var key in imageSources) {
      images[key] = new Image();
      images[key].src = imageSources[key];

      images[key].onload = function() {
        loadedImages++;
        percentage = Math.floor(loadedImages * 100 / totalImages);

        view.setHTML(loadingPercentage, percentage + '%'); //displaying percentage

        if (loadedImages >= totalImages) {
          view.removeFromBody(loadingPercentage);
          that.initMainApp();
        }
      };
    }
  };

  this.initMainApp = function() {
    var marioMakerInstance = MarioMaker.getInstance();
    marioMakerInstance.init();
  };
}

window.onload = function() {
  var preloader = new Preloader();
  preloader.init();
};
