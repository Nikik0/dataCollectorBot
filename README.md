# dataCollectorBot

		/*
			флоу
			запрос на обработку персональных данных
			запрос имени
			запрос фамилии
			запрос даты рождения
			запрос почты
			подтверждение (меню все ок или изменить, если изменить то меню что изменить с кнопками имя-почта)
			после подтверждения либо коммитим все либо возможность изменить (уточнить требования)
			присылаем картинку что все клево и говорим что регистрация прошла успешно
		*/

		/*

			состояние храним в стейт машине, соответственно изменять данные можем через переход стейта на прошлое состояние
			фактически стейтмашина это интерфейс с двумя полями: стейт и флаг есть ли след стейт
			сам стейт имеет метод совершения действия (обновление поля сущности пользователя полученными данными, начало работы, подтверждение, завершение взаимодействия)
		*/