package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	runningTrainingType  = "Бег"
	walkingTrainingType  = "Ходьба"
	fieldsInTrainingData = 3
)

// Training содержит данные о тренировке и персональные данные пользователя.
type Training struct {
	Steps        int           // количество шагов за тренировку.
	TrainingType string        // тип тренировки (бег или ходьба).
	Duration     time.Duration // длительность тренировки.
	personaldata.Personal
}

// Parse разбирает строку формата "3456,Ходьба,3h00m"
// и записывает полученные данные в структуру Training.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != fieldsInTrainingData {
		return fmt.Errorf("неверный формат данных тренировки: %q", datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	trainingType := parts[1]
	if trainingType == "" {
		return errors.New("тип тренировки не может быть пустым")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("не удалось разобрать продолжительность тренировки: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность тренировки должна быть положительной")
	}

	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var (
		calories float64
		err      error
	)

	switch t.TrainingType {
	case runningTrainingType:
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case walkingTrainingType:
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories,
	)

	return info, nil
}
