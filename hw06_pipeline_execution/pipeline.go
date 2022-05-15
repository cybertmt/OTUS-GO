package main

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

// StageControl реализует ранний выход из пайплайна.
func StageControl(done In, in In) Out {
	// Открываем возвращаемый канал элементов.
	out := make(Bi)
	// В горутине слушаем каналы done и in.
	go func() {
		defer close(out)
		for {
			select {
			case <-done: // Отсанавливаем пайплайн по сигналу.
				return
			case v, ok := <-in: // Перекладываем элементы из in в out.
				{
					if !ok { // Элементы в канале закончились, выходим.
						return
					}
					out <- v
				}
			}
		}
	}()
	// Возвращаем канал элементы для обработки stage.
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Элементы канала in проходят этапы пайплайна.
	// StageControl реализует ранний выход из функции (<-done).
	for _, stage := range stages {
		in = stage(StageControl(done, in))
	}
	// Возвращаем канал с элементами, прощедшими список этапов (stages).
	return in
}
