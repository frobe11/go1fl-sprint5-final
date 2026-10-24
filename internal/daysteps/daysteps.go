package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const fieldsInDayStepsData = 2

// DaySteps содержит данные о дневной прогулке и персональные данные пользователя.
type DaySteps struct {
	Steps    int           // количество шагов за прогулку.
	Duration time.Duration // длительность прогулки.
	personaldata.Personal
}

// Parse разбирает строку формата "678,0h50m"
// и записывает полученные данные в структуру DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != fieldsInDayStepsData {
		return fmt.Errorf("неверный формат данных прогулки: %q", datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("не удалось разобрать продолжительность прогулки: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность прогулки должна быть положительной")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	)

	return info, nil
}
