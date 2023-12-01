## tesgode
**Tes**t for **Go** **de**veloper. Небольшая, в меру интересная, с ожиданием креатива и ответственности - задачка для Go-разработчика :-)

##### Задача
Написать сервис анализа нарушений ПДД.
В целях уменьшения временных и трудовых затрат ожидается минимальный функционал, обеспечивающий правильную работу сервиса. Нарушением ПДД в рамках данной задачи считается только проезд на красный свет.
Для работы сервиса ожидается использование брокера сообщений [cat](https://github.com/kvolis/tesgode/tree/main/cat) и базы данных [dog](https://github.com/kvolis/tesgode/tree/main/dog). В них представлены простейшие симуляции.

Тестируемому кандидату необходимо:
- ознакомиться с документацией пакетов брокера `Cat` и базы данных `Dog`
- самостоятельно продумать архитектуру приложения
- определить структуру проекта
- обеспечить надёжное и бесперебойное функционирование, обрабатывать ошибки
- правильно завершить работу сервиса
- использовать многопоточность при работе, Cat и Dog специально имитируют задержки выполнения работы

Условие завершения работы приложния остаётся на усмотрение тестируемого. Рекомендуется "принять" от брокера и "сохранить" в БД не менее 100 сообщений.

##### Описание

При старте сервиса инициализировать конфиг (его представить в виде структуры без полей). Реализация методов не требуется, но возможна на усмотрение тестируемого, как и структура конфига.

Далее ожидается подключение к брокеру и БД. Брокер при этом вернёт канал, в который начнут идти сообщения.

Сообщение - это слайс байтов JSON-представления объекта типа `Passage` из пакета [models](github.com/kvolis/tesgode/tree/main/models). `Passage` - информация о проезде транспортного средства с гос.номером `LicenseNum` и набором точек `Track` (таймстемп является представлением времени Unix в секундах).
Нужно проанализировать информацию на наличие нарушения - проезда на красный свет.
_Светофор горит зелёным в любую минуту с 0-й по 44-ю секунду включительно, а с 45-й по 59-ю включительно - красным. Например, в 13:38:26 горит зелёный, а в 20:07:51 - красный._
_Важно!_ Только последняя по времени точка попадает за линию светофора, поэтому анализировать необходимо только её. Обратите внимание на формулировку "последняя по времени".

Если нарушения ПДД нет, перейти к следующему проезду, никаких действий не требуется (но остаётся на усмотрение тестируемого).

При наличии нарушения нужно сохранить о нём информацию в БД. Данные для сохранения остаются на усмотрение тестируемого. БД представляет из себя имитацию таблицы с полями ID (int), Key (string) и Value (string). В случае успешного "сохранения" будет возвращён ID новой записи.

После обработки примерно 100 сообщений нужно правильно завершить работу приложения.

##### Советы и пожелания
- база данных работает с каждым днём всё хуже, поэтому через некоторое время Компании придётся перейти на другую; попробуйте предусмотреть эту возможность
- в любой непонятной ситуации действуйте на своё усмотрение; цель данной задачи - не жёсткое следование каким-либо правилам, а определение способности рассуждать, анализировать, креативить
- ожидаемое количество кода - около 100 строк, но не нужно стремиться именно к этому числу
- задача сформулирована в расчёте на выполнение в течение 1 часа, не считая изучения документации (абсолютно нормально, если получится больше или меньше)
- если у вас появились вопросы или пожелания относительно реализаций Cat или Dog, зафиксируйте их для дальнейшего обсуждения
- будут приветствоваться любые вольности, при этом основная задача должна быть выполнена
- постарайтесь понять, почему в каком-то конкретном месте вы поступили именно так, чтобы потом донести это до нас и убедить при необходимости :-)
- удачи!


