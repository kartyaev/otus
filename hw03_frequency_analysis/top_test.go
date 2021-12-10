package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var additionalText = `Белеет парус одинокой
					В тумане моря голубом!..
					Что ищет он в стране далекой?
					Что кинул он в краю родном?..
					
					Играют волны — ветер свищет,
					И мачта гнется и скрыпит…
					Увы! он счастия не ищет
					И не от счастия бежит!
					
					Под ним струя светлей лазури,
					Над ним луч солнца золотой…
					А он, мятежный, просит бури,
					Как будто в бурях есть покой!`

//nolint
var linusEmail = `Hello everybody out there using minix - 
				
				I'm doing a (free) operating system (just a hobby, won't be big and 
				professional like gnu) for 386(486) AT clones. This has been brewing 
				since april, and is starting to get ready. I'd like any feedback on 
				things people like/dislike in minix, as my OS resembles it somewhat 
				(same physical layout of the file-system (due to practical reasons) 
				among other things). 
				
				I've currently ported bash(1.08) and gcc(1.40), and things seem to work. 
				This implies that I'll get something practical within a few months, and 
				I'd like to know what features most people would want. Any suggestions 
				are welcome, but I won't promise I'll implement them :-) 
				
				Linus (torv...@kruuna.helsinki.fi) 
				
				PS. Yes - it's free of any minix code, and it has a multi-threaded fs. 
				It is NOT protable (uses 386 task switching etc), and it probably never 
				will support anything other than AT-harddisks, as that's all I have :-(. `

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})

	t.Run("positive test additional text", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"в",       // 4
				"он",      // 4
				"и",       // 3
				"ищет",    // 2
				"не",      // 2
				"ним",     // 2
				"счастия", // 2
				"что",     // 2
				"а",       // 1
				"бежит",   // 1
			}

			require.Equal(t, expected, Top10(additionalText))
		}
	})

	t.Run("positive test linus email", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"and",    // 7
				"a",      // 4
				"it",     // 4
				"to",     // 5
				"any",    // 3
				"like",   // 3
				"minix",  // 3
				"things", // 3
				"as",     // 2
				"free",   // 2
			}

			require.Equal(t, expected, Top10(linusEmail))
		}
	})
}
