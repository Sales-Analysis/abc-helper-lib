# Inventory Analyses

This page summarizes the inventory-focused analyses currently implemented in
`abc-helper-lib`.

## Summary Table

| Analysis | Package | Facade Method | Primary Inputs | Main Output |
| :--- | :--- | :--- | :--- | :--- |
| `ABC` | `abc` | `Inventory().ABC` | quantity, unit price | revenue share and group |
| `XYZ` | `xyz` | `Inventory().XYZ` | demand series | demand variability group |
| `ABC-XYZ` | `abcxyz` | `Inventory().ABCXYZ` | sales value and demand series | combined group |
| `VED` | `ved` | `Inventory().VED` | criticality score | criticality group |
| `FSN` | `fsn` | `Inventory().FSN` | movement history | fast/slow/non-moving group |
| `HML` | `hml` | `Inventory().HML` | unit cost | cost-based group |
| `SDE` | `sde` | `Inventory().SDE` | lead time in days | procurement difficulty group |
| `EOQ` | `eoq` | `Inventory().EOQ` | demand, ordering cost, holding cost | optimal order quantity |
| `Reorder Point` | `reorderpoint` | `Inventory().ReorderPoint` | average demand, lead time, safety stock | reorder point |
| `Safety Stock` | `safetystock` | `Inventory().SafetyStock` | demand deviation, lead time, service factor | safety stock |
| `Pareto 80/20` | `pareto` | `Inventory().Pareto` | value per item | cumulative value and 80/20 check |
| `GM/Contribution` | `gmcontribution` | `Inventory().GMContribution` | revenue and cost components | margin metrics |
| `Service Level` | `servicelevel` | `Inventory().ServiceLevel` | demand fulfillment and stockout cycles | service level metrics |

## ABC

- Purpose: classify items by revenue contribution.
- Input: `Quantity`, `Price`, optional thresholds.
- Calculation:
  `PriceTotal = Quantity * Price`,
  `ShareTotal = PriceTotal / TotalRevenue * 100`,
  `ShareAccumulated` is cumulative after sorting by `PriceTotal` descending.
- Default thresholds:
  `A <= 80%`, `B <= 95%`, `C > 95%`.
- Output includes `PriceTotal`, `ShareTotal`, `ShareAccumulated`, and `Group`.

## XYZ

- Purpose: classify demand by stability.
- Input: a `Demands` series per item and optional CV thresholds.
- Calculation:
  mean demand, standard deviation, coefficient of variation.
- Default thresholds:
  `X <= 10`, `Y <= 25`, otherwise `Z`.
- Special case: zero mean demand or an empty demand series is classified as `Z`.

## ABC-XYZ

- Purpose: combine revenue importance and demand stability.
- Input: item value inputs for `ABC` plus demand history for `XYZ`.
- Composition:
  runs `ABC` and `XYZ` independently and merges both results by original item.
- Output includes `ABCGroup`, `XYZGroup`, and `CombinedGroup` such as `AX`.

## VED

- Purpose: classify items by operational criticality.
- Input: `CriticalityScore`.
- Default thresholds:
  `V >= 70`, `E >= 40`, otherwise `D`.
- Output contains the original score and the resulting `V/E/D` group.

## FSN

- Purpose: classify movement frequency.
- Input: `Movements` across periods.
- Calculation:
  active periods, activity rate, total movement, average movement, last movement period.
- Default thresholds:
  `F >= 70%` activity,
  `S >= 1%` activity,
  `N` when there are no active periods or activity is below the slow threshold.

## HML

- Purpose: classify items by unit cost.
- Input: `UnitCost`, optional explicit thresholds.
- If thresholds are not provided, the package derives them from the data by
  splitting sorted values into three bands.
- Output groups items into `H`, `M`, or `L`.

## SDE

- Purpose: classify procurement difficulty by lead time.
- Input: `LeadTimeDays`, optional explicit thresholds.
- If thresholds are not provided, the package derives them from the sorted lead
  time distribution.
- Output groups items into `S`, `D`, or `E`.

## EOQ

- Purpose: calculate the economic order quantity.
- Input: `AnnualDemand`, `OrderingCost`, `HoldingCost`.
- Formula:
  `EOQ = sqrt((2 * AnnualDemand * OrderingCost) / HoldingCost)`.
- Additional output:
  `OrdersPerYear = AnnualDemand / EOQ`.
- Any non-positive parameter produces zero output for the affected metric.

## Reorder Point

- Purpose: calculate when replenishment should start.
- Input: `AverageDemandPerPeriod`, `LeadTimePeriods`, `SafetyStock`.
- Formula:
  `LeadTimeDemand = AverageDemandPerPeriod * LeadTimePeriods`,
  `ReorderPoint = LeadTimeDemand + SafetyStock`.

## Safety Stock

- Purpose: estimate buffer stock against demand variability.
- Input: `DemandStdDev`, `LeadTimePeriods`, optional `ServiceFactor`.
- Formula:
  `SafetyStock = ServiceFactor * DemandStdDev * sqrt(LeadTimePeriods)`.
- Default `ServiceFactor`: `1.65`.

## Pareto 80/20

- Purpose: check whether a small share of items drives most of the value.
- Input: item `Value`, optional thresholds for top value share and top item share.
- Default thresholds: `80%` value and `20%` items.
- Output:
  per-item value shares and a summary showing whether the configured Pareto
  target is met.

## GM/Contribution

- Purpose: calculate gross margin and contribution metrics.
- Input: `Revenue`, `COGS`, `VariableCost`, `FixedCost`.
- Output includes:
  `GrossMargin`, `GrossMarginRate`,
  `ContributionMargin`, `ContributionRate`,
  `NetContribution`.

## Service Level

- Purpose: estimate fulfillment quality.
- Input:
  `DemandedUnits`, `FulfilledUnits`, `TotalCycles`, `StockoutCycles`.
- Calculated metrics:
  `FillRate`,
  `CycleServiceLevel`,
  `StockoutRate`,
  `UnfulfilledUnits`.
- Selection rule:
  if `TotalCycles > 0`, the exported `ServiceLevel` uses `CycleServiceLevel`;
  otherwise it falls back to `FillRate`.
