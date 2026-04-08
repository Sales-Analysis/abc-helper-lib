# Товарные анализы

На этой странице собраны товарные и supply-chain анализы, которые сейчас
реализованы в `abc-helper-lib`.

## Сводная таблица

| Анализ | Пакет | Facade-метод | Основные входы | Ключевой результат |
| :--- | :--- | :--- | :--- | :--- |
| `ABC` | `abc` | `Inventory().ABC` | количество, цена | доля выручки и группа |
| `XYZ` | `xyz` | `Inventory().XYZ` | ряд спроса | группа по вариативности |
| `ABC-XYZ` | `abcxyz` | `Inventory().ABCXYZ` | стоимость продаж и ряд спроса | комбинированная группа |
| `VED` | `ved` | `Inventory().VED` | score критичности | группа критичности |
| `FSN` | `fsn` | `Inventory().FSN` | история движений | fast/slow/non-moving группа |
| `HML` | `hml` | `Inventory().HML` | unit cost | группа по стоимости |
| `SDE` | `sde` | `Inventory().SDE` | lead time в днях | группа по сложности снабжения |
| `EOQ` | `eoq` | `Inventory().EOQ` | спрос, ordering cost, holding cost | оптимальный размер заказа |
| `Reorder Point` | `reorderpoint` | `Inventory().ReorderPoint` | средний спрос, lead time, safety stock | точка заказа |
| `Safety Stock` | `safetystock` | `Inventory().SafetyStock` | отклонение спроса, lead time, service factor | страховой запас |
| `Pareto 80/20` | `pareto` | `Inventory().Pareto` | value по товару | накопленная доля и проверка 80/20 |
| `GM/Contribution` | `gmcontribution` | `Inventory().GMContribution` | выручка и компоненты затрат | маржинальные метрики |
| `Service Level` | `servicelevel` | `Inventory().ServiceLevel` | исполнение спроса и stockout cycles | метрики уровня сервиса |

## ABC

- Назначение: классификация товаров по вкладу в выручку.
- Вход:
  `Quantity`, `Price`, опционально свои пороги.
- Расчет:
  `PriceTotal = Quantity * Price`,
  `ShareTotal = PriceTotal / TotalRevenue * 100`,
  `ShareAccumulated` считается накопительно после сортировки по `PriceTotal`
  по убыванию.
- Дефолтные пороги:
  `A <= 80%`, `B <= 95%`, `C > 95%`.
- Результат:
  `PriceTotal`, `ShareTotal`, `ShareAccumulated`, `Group`.

## XYZ

- Назначение: классификация по стабильности спроса.
- Вход:
  ряд `Demands` по каждому товару и опциональные пороги CV.
- Расчет:
  средний спрос, стандартное отклонение и коэффициент вариации.
- Дефолтные пороги:
  `X <= 10`, `Y <= 25`, иначе `Z`.
- Спецслучай:
  нулевой средний спрос или пустой ряд спроса попадает в `Z`.

## ABC-XYZ

- Назначение: объединить важность по выручке и стабильность спроса.
- Вход:
  входы для `ABC` плюс история спроса для `XYZ`.
- Композиция:
  пакет независимо запускает `ABC` и `XYZ`, а затем объединяет результаты по
  исходному индексу товара.
- Результат:
  `ABCGroup`, `XYZGroup`, `CombinedGroup`, например `AX`.

## VED

- Назначение: классификация по критичности.
- Вход:
  `CriticalityScore`.
- Дефолтные пороги:
  `V >= 70`, `E >= 40`, иначе `D`.
- Результат:
  исходный score и группа `V/E/D`.

## FSN

- Назначение: классификация по частоте движения.
- Вход:
  массив `Movements` по периодам.
- Расчет:
  число активных периодов, activity rate, total movement, average movement,
  номер последнего периода с движением.
- Дефолтные пороги:
  `F >= 70%` активности,
  `S >= 1%` активности,
  `N`, если активных периодов нет или активность ниже slow-порога.

## HML

- Назначение: классификация по unit cost.
- Вход:
  `UnitCost`, опционально свои пороги.
- Если пороги не заданы, пакет выводит их из данных, разбивая отсортированные
  значения на три диапазона.
- Результат:
  группа `H`, `M` или `L`.

## SDE

- Назначение: классификация по сложности снабжения через lead time.
- Вход:
  `LeadTimeDays`, опционально свои пороги.
- Если пороги не заданы, пакет выводит их из распределения lead time.
- Результат:
  группа `S`, `D` или `E`.

## EOQ

- Назначение: расчет экономически оптимального размера заказа.
- Вход:
  `AnnualDemand`, `OrderingCost`, `HoldingCost`.
- Формула:
  `EOQ = sqrt((2 * AnnualDemand * OrderingCost) / HoldingCost)`.
- Дополнительный результат:
  `OrdersPerYear = AnnualDemand / EOQ`.
- Если любой из ключевых параметров неположительный, метрика возвращается как
  `0`.

## Reorder Point

- Назначение: определить момент старта пополнения.
- Вход:
  `AverageDemandPerPeriod`, `LeadTimePeriods`, `SafetyStock`.
- Формула:
  `LeadTimeDemand = AverageDemandPerPeriod * LeadTimePeriods`,
  `ReorderPoint = LeadTimeDemand + SafetyStock`.

## Safety Stock

- Назначение: оценить буферный запас против вариативности спроса.
- Вход:
  `DemandStdDev`, `LeadTimePeriods`, опционально `ServiceFactor`.
- Формула:
  `SafetyStock = ServiceFactor * DemandStdDev * sqrt(LeadTimePeriods)`.
- Дефолтный `ServiceFactor`: `1.65`.

## Pareto 80/20

- Назначение: проверить, дает ли малая доля товаров основную долю value.
- Вход:
  `Value` по товару, опционально свои пороги top value share и top item share.
- Дефолтные пороги:
  `80%` value и `20%` товаров.
- Результат:
  доли по каждому товару и summary, показывающий, выполняется ли заданный
  критерий Pareto.

## GM/Contribution

- Назначение: расчет gross margin и contribution метрик.
- Вход:
  `Revenue`, `COGS`, `VariableCost`, `FixedCost`.
- Результат включает:
  `GrossMargin`, `GrossMarginRate`,
  `ContributionMargin`, `ContributionRate`,
  `NetContribution`.

## Service Level

- Назначение: оценка качества исполнения спроса.
- Вход:
  `DemandedUnits`, `FulfilledUnits`, `TotalCycles`, `StockoutCycles`.
- Вычисляемые метрики:
  `FillRate`,
  `CycleServiceLevel`,
  `StockoutRate`,
  `UnfulfilledUnits`.
- Правило выбора:
  если `TotalCycles > 0`, итоговый `ServiceLevel` берется из
  `CycleServiceLevel`, иначе используется `FillRate`.
