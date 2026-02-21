package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/disk"
)

func main() {
	
	const cpuThreshold = 80.0  // CPU > 80% считается критическим
	const memThreshold = 80.0  // Память > 80% критично
	const diskThreshold = 90.0 // Диск > 90% критично

	fmt.Println("🖥️  Сервер мониторинг запущен...")

	for {
		
		cpuPercent, _ := cpu.Percent(0, false) // получаем среднюю загрузку всех ядер
		if len(cpuPercent) > 0 && cpuPercent[0] > cpuThreshold {
			fmt.Printf("⚠️  Внимание! CPU загрузка критическая: %.2f%%\n", cpuPercent[0])
		}

		
		vmStat, _ := mem.VirtualMemory()
		if vmStat.UsedPercent > memThreshold {
			fmt.Printf("⚠️  Внимание! Использование памяти критическое: %.2f%%\n", vmStat.UsedPercent)
		}

		
		diskStat, _ := disk.Usage("/")
		if diskStat.UsedPercent > diskThreshold {
			fmt.Printf("⚠️  Внимание! Использование диска критическое: %.2f%%\n", diskStat.UsedPercent)
		}

		
		fmt.Printf("CPU: %.2f%%, Memory: %.2f%%, Disk: %.2f%%\n", cpuPercent[0], vmStat.UsedPercent, diskStat.UsedPercent)

		time.Sleep(5 * time.Second) // проверяем каждые 5 секунд
	}
}