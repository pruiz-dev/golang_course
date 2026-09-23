# CURRICULUM V2: Middle Golang Developer

Этот план направлен на глубокое понимание внутреннего устройства Go и интенсивную практику. Каждая тема может содержать множество практических задач (циклы `/[lesson]` и `/[practice]`), пока концепция не будет усвоена на 100%.

## Этап 1: Memory & Foundation
- [ ] `01_memory_and_types` (Выравнивание, Padding, архитектурные типы `int` vs `int32`, `uintptr`)
- [ ] `02_pointers_and_references` (Pass by value, escape analysis, Heap vs Stack)
- [ ] `03_advanced_functions` (Замыкания, отложенные вызовы `defer`, паники и восстановление)

## Этап 2: Data Structures Deep Dive
- [ ] `04_slice_internals` (Slice header, capacity growth formula, memory leaks with slicing)
- [ ] `05_map_internals` (hmap, bmap, hashing, load factor, evacuation, concurrency issues)
- [ ] `06_string_internals` (String header, runes vs bytes, allocation-free conversions)

## Этап 3: Concurrency Primitives
- [ ] `07_goroutines` (GMP Scheduler, мертвые блокировки, утечки горутин)
- [ ] `08_channels` (Буферизированные vs небуферизированные, `select`, закрытие каналов)
- [ ] `09_context` (Отмена горутин, таймауты, передача значений)

## Этап 4: Synchronization
- [ ] `10_mutex` (Mutex vs RWMutex, race conditions, go run -race)
- [ ] `11_waitgroup_and_pool` (sync.WaitGroup, sync.Pool для экономии аллокаций)
- [ ] `12_atomic` (Атомарные операции vs Mutex)

## Этап 5: Architecture & Interfaces
- [ ] `13_interfaces` (Пустой интерфейс, внутреннее устройство интерфейса: itab и data)
- [ ] `14_solid` (Принципы SOLID в реалиях Go)
- [ ] `15_error_handling` (Обертывание ошибок, пользовательские типы ошибок)

## Этап 6: Проекты (Портфолио)
- [ ] `project_01_cli_tool` (Разработка сложной CLI-утилиты с конкурентной обработкой файлов)
- [ ] `project_02_cache_server` (Собственный in-memory кэш сервер по TCP с инвалидацией и TTL)
- [ ] `project_03_microservice` (REST API микросервис с базой данных и чистой архитектурой)
