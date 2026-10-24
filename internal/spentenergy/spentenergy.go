package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Distance рассчитывает пройденную дистанцию в километрах
// по количеству шагов и росту пользователя в метрах.
func Distance(steps int, height float64) float64 {
	// Отрицательное количество шагов или некорректный рост
	// не имеют физического смысла - возвращаем нулевую дистанцию.
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}

// MeanSpeed рассчитывает среднюю скорость в км/ч
// по количеству шагов, росту пользователя и продолжительности активности.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Без положительной продолжительности средняя скорость не определена.
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

// validateParameters проверяет, что все входные параметры имеют смысл
// для расчёта потраченных калорий.
func validateParameters(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}

	return nil
}

// RunningSpentCalories рассчитывает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateParameters(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

// WalkingSpentCalories рассчитывает количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateParameters(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH
	calories *= walkingCaloriesCoefficient

	return calories, nil
}
