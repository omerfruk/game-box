
var debug = '';
$(function(){
	var util = {
		randomize: function(from, to, r){
			return Math.floor( Math.random() * (to - from) + from );
		},
		shuffle: function(v){
			for(var j, x, i = v.length; i; j = parseInt(Math.random() * i), x = v[--i], v[i] = v[j], v[j] = x);
			return v;
		},
		local: {
			set:function(key,val){
				if(Modernizr.localstorage){
					localStorage[key] = val;
				}
			},
			get: function(key){
				if(Modernizr.localstorage){
					return localStorage[key];
				} else return null;
			}
		}
	};
	var suit = ['spades','hearts','diams','clubs'];
	var cards = ['A','2','3','4','5','6','7','8','9','10','J','Q','K'];
	var score = [11,2,3,4,5,6,7,8,9,10,10,10,10];
	var deck = [];
	var game = [];
	var speed = 500;
	var playerHand = $('.player-hand');
	var casino = $('.casino');
	var hitstand = $('.hit-btn, .stand-btn');
	var dealbtn = $('.deal-btn .btn, .bet-val');
	var account = $('.bank span');
	var total = {
		player:0,
		casino:0
	};
	var status = 'idle';
	var stake = 0;
	var money = util.local.get('21money') ? Math.max(util.local.get('21money'),100): 100;

	account.html(money);
	var deal = function(end){
		if(end === true){
			if(total.casino == total.player){
				say('Kal. '+total.casino);
				acc.give(+stake);
			} else if(total.casino<total.player){
				say('Sen Kazandın $'+(stake*2) );
				acc.give(stake*2);
			} else {
				say(total.casino+'. Casino Kazandı');
			}
			replay();
		}
		if(game.length){
			var turn = game.shift();
			var cardData = deck.shift();
			debug += JSON.stringify(cardData)+',';
			/*card html*/
			var card = '';
			var mark = '<span class="mark">\
							<h1>'+cardData.value+'</h1>\
							<h2>&'+cardData.suit+';</h2>\
						</span>';
			var suits = '<span class="suit">&'+cardData.suit+';</span>';
			if(cardData.order<10){
				for(var l=0; l<cardData.order; l++){
					suits += '<span class="suit">&'+cardData.suit+';</span>';
				}
			}
			card += '<div class="card '+(turn.player? 'player' : 'casino')+'-card val'+cardData.value+' '+cardData.suit+' anim"><div class="face">';
			card += mark+suits+mark;
			card += '</div><div class="back"></div></div>';
			card = $(card);
			card.css({top:$('.deck').offset().top, left:$('.deck').offset().left});
			var pack = casino;
			if(turn.player) {
				pack = playerHand;
				total.player += cardData.score;
			} else {
				total.casino += cardData.score;
			}

			/*card animation*/
			var stack = pack.find('.card').length;
			var size = pack.width()*.5;
			$('body').append(card);
			card.animate({top:pack.offset().top},speed,function(){
				$(this).remove();
				pack.append(card);
				card.removeClass('anim').attr('style', Modernizr._prefixes.join('transform:rotate(' + util.randomize(-4,3)+'deg);')).css({left:size-55});
				var position = size-((stack+1)*55);
				for(var i = 0; i<=stack; i++){
					var c = pack.find('.card').eq(i);
					var nextPos = position + util.randomize(105,115)*i;
					c.animate({left:nextPos});
				}
			});
			if(turn.flip) card.addClass('flip');
			if(game.length){
				setTimeout(deal,speed);
			}

			if(turn.player){
				if(total.player>21){
					var ace = $('.valA.player-card');
					if(ace.length>0 && $('.player-card.valAone').length===0){
						total.player -= 10;
						ace.toggleClass('valA valAone');
						if (total.player === 21 ){
							say('Black Jack! Sen Kazandın '+(stake*3)+'$');
							acc.give(stake*3);
							replay();
						} else {
							say('Senin puanın '+total.player);
						}
					} else {
						say('Battın');
						hitstand.disable();
						replay();
					}
				} else if (total.player === 21 ){
					say('Black Jack! Sen kazandın '+(stake*3)+'$');
					hitstand.disable();
					acc.give(stake*3);
					replay();
				} else {
					say('Senin puanın '+total.player);
				}
				if(pack.find('.card').length>0 && status !== 'end' ){
					setTimeout(function(){
						hitstand.enable();
					},1000);
				}
			} else if(status == 'casinoTurn'){
				if(total.casino>21){
					var ace = $('.valA.casino-card');
					if(ace.length>0 && $('.casino-card.valAone').length===0){
						total.casino -= 10;
						ace.toggleClass('valA valAone');
						if(total.casino===21){
							say('Black Jack. Casino kazandı');
							replay();
						} else {
							if(status === 'casinoTurn'){
								setTimeout(casinoDeal,speed);
							}
						}
					} else {
						acc.give(stake*2);
						replay();
						say('Casino battı. Sen kazandın '+stake+'$');
						acc.give(stake*2);
					}
				} else if(total.casino===21){
					say('Black Jack. Casino kazandı');
					replay();
				} else {
					if(status === 'casinoTurn'){
						setTimeout(casinoDeal,speed);
					}
				}
			}
		}
	};
	var replay = function(){
		status = 'end';
		setTimeout(redeal,3000);
	};

	var dealer = function(){
		deck = [];
		for(var i=0; i<4; i++){
			for(var l=0; l<13; l++){
				deck.push({value:cards[l],suit:suit[i],order:l, score:score[l]});
			}
		}

		game = [
			{player:true, flip:true},
			{player:false, flip:true},
			{player:true, flip:true},
			{player:false, flip:false}
		];
		total = {
			player:0,
			casino:0
		};
		deck = util.shuffle(deck);

		say('Birdaha');
		debug += '<<NEW GAME>>';
		stake = $('.bet-val').val();
		if(stake<5 || stake>100){
			stake = 10;
			$('.bet-val').val(10);
		}
		if(stake>money){
			stake = money;
			$('.bet-val').val(stake);
		}
		util.local.set('21stake',stake);
		status = 'game';
		acc.take(stake);
		deal();
	};

	var redeal = function(){
		status = 'idle';
		var center = $('.player-hand').width()*.5;
		$('.player-hand .card, .casino .card').animate({left:center-55, opacity:0},speed,function(){
			$(this).remove();
		});

		if(money===0){
			$('.autodeal').prop('checked',false);
			say('Senin için bir abimmiz 10$ bıraktı😉 Bunu iyi kullan!','casino',5000);
			acc.give(10);
			$('.bet-val').val(5);
		}

		if($('.autodeal').prop('checked')){
			setTimeout(dealer,speed*1.2);
		} else {
			dealbtn.enable();
		}
	};

	var casinoDeal = function(){
		say(total.casino);
		if(total.casino>=17 || total.casino>total.player){
			setTimeout(function(){
				deal(true);
			},speed);
		} else {
			game.push({player:false, flip:true});
			setTimeout(deal,speed);
		}
	};

	var say = function(what, who, display) {
		var target = who||'casino';
		var balloon = $('<div class="speakballon '+target+'-speak">'+what+'</div>');
		setTimeout(function(){
			$('.'+target+'-speak').remove();
			balloon.appendTo('body').fadeIn();
		},speed);
		setTimeout(function(){
			balloon.fadeOut(function(){
				$(this).remove();
			});
		},(display||2000));
	};

	var acc = {
		take:function(i){
			money -= i;
			account.html(money);
			util.local.set('21money',money);
		},
		give:function(i){
			money += i;
			account.html(money);
			util.local.set('21money',money);
		}
	};

	$('.hit-btn').on('click', function(){
		debug += '<<HIT>>';
		hitstand.disable();
		game.push({player:true, flip:true});
		say('Kart ver!', 'player');
		deal();
	});
	$('.stand-btn').on('click', function(){
		debug += '<<STAND>>';
		hitstand.disable();
		status = 'casinoTurn';
		$('.casino .card:eq(1)').addClass('flip');
		say('Kal','player');
		setTimeout(casinoDeal,speed);
	});

	$('.autodeal').on('click',function(){
		dealbtn.disable();
		var btn = $('.deal-btn .btn');
		var isOn = btn.toggleClass('btn-success').hasClass('btn-success');
		if(status==='idle'){
			dealer();
		}
	});
	$('.deal-btn .btn').on('click', function(){
		dealbtn.disable();
		dealer();
	});

	say('Paranı Koy');
	$(document).on('keydown',function(e){
		switch (e.keyCode){
			case 68:
				$('.deal-btn .btn:enabled').click();
				break;
			case 72:
				$('.hit-btn:enabled').click();
				break;
			case 83:
				$('.stand-btn:enabled').click();
				break;
			case 38:
				$('.bet-val').val(Math.min(+$('.bet-val').val()+5, 100));
				break;
			case 40:
				$('.bet-val').val(Math.max(+$('.bet-val').val()-5, 5));
				break;
			case 65:
				$('.autodeal').click();
				break;
		}
	});
	;(function($) {

		$.fn.extend({
			enable: function() {
				this.prop('disabled',false);
				return this;
			},
			disable: function() {
				this.prop('disabled',true);
				return this;
			}
		});
	})(jQuery);
	$('.bet-val').val(util.local.get('21stake')|| 10);
});	